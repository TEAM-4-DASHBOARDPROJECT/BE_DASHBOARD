package delivery

import (
	"immersiveProject/features/users"
	"immersiveProject/middlewares"
	"immersiveProject/utils/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	data users.ServiceInterface
}

func New(e *echo.Echo, usecase users.ServiceInterface) {

	handler := UserHandler{
		data: usecase,
	}

	e.GET("/users", handler.GetAllWithJWT, middlewares.JWTMiddleware())
	e.GET("/users/:id", handler.GetByIdWithJWT, middlewares.JWTMiddleware())
	e.POST("/users", handler.AddUser, middlewares.JWTMiddleware())
	e.PUT("/users/:id", handler.PutDataWithJWT, middlewares.JWTMiddleware())
	e.DELETE("/users/:id", handler.DeldateWithJWT, middlewares.JWTMiddleware())

}

func (users *UserHandler) GetAllWithJWT(e echo.Context) error {

	idToken, _ := middlewares.ExtractToken(e)
	res, err := users.data.GetAll(idToken)
	if err != nil {
		return e.JSON(400, helper.FailedResponseHelper("err"))
	}

	respon := toResponList(res)

	return e.JSON(200, helper.SuccessDataResponseHelper("succes get all data", respon))

}

func (users *UserHandler) GetByIdWithJWT(e echo.Context) error {

	idToken, _ := middlewares.ExtractToken(e)
	id := helper.ParamInt(e)
	if id == -1 {
		return e.JSON(400, helper.FailedResponseHelper("param must be number"))
	}

	res, err := users.data.GetById(id, idToken)
	if err != nil {
		return e.JSON(400, helper.FailedResponseHelper("id not found"))
	}

	respon := toResponId(res)

	return e.JSON(200, helper.SuccessDataResponseHelper("succes get data by id", respon))

}

func (users *UserHandler) AddUser(e echo.Context) error {

	var req UserReq
	err := e.Bind(&req)
	if err != nil {
		return e.JSON(400, helper.FailedResponseHelper("err"))
	}
	idToken, _ := middlewares.ExtractToken(e) // barusan di tambah
	// id := helper.ParamInt(e)
	if idToken != 1 {
		return e.JSON(400, helper.FailedResponseHelper("kamu tidak bukan superman"))
	}

	add := ToCore(req)
	row, _ := users.data.PostData(add)
	if row == 1 {
		return e.JSON(200, helper.SuccessResponseHelper("succes insert data"))
	} else {
		return e.JSON(400, helper.FailedResponseHelper("errPost"))
	}

}

func (users *UserHandler) PutDataWithJWT(e echo.Context) error {

	id := e.Param("id")
	idProd, _ := strconv.Atoi(id)
	idFromToken, _ := middlewares.ExtractToken(e)
	prodReq := UserReq{}
	err := e.Bind(&prodReq)

	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.FailedResponseHelper("failed to bind data, check your input"))
	}
	dataProduct := ToCore(prodReq)
	row, errUpd := users.data.PutData(idProd, idFromToken, dataProduct)
	if errUpd != nil {
		return e.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("you dont have access"))
	}
	if row == 0 {
		return e.JSON(http.StatusBadRequest, helper.FailedResponseHelper("failed to update data"))
	}
	return e.JSON(http.StatusOK, helper.SuccessResponseHelper("success"))

}

func (users *UserHandler) DeldateWithJWT(e echo.Context) error {

	idToken, _ := middlewares.ExtractToken(e)
	id := helper.ParamInt(e)
	if id == -1 {
		return e.JSON(400, helper.FailedResponseHelper("param must be number"))
	}

	row, _ := users.data.DeleteData(id, idToken)
	if row == 1 {
		return e.JSON(200, helper.SuccessResponseHelper("succes delete data"))
	} else {
		return e.JSON(400, helper.FailedResponseHelper("not have access"))
	}

}
