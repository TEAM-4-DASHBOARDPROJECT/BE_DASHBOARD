package delivery

import (
	"immersiveProject/features/users"
	"immersiveProject/middlewares"
	"immersiveProject/utils/helper"
	"strconv"

	"github.com/labstack/echo/v4"
)

type userDelivery struct {
	userUsecase users.UsecaseInterface
}

func New(e *echo.Echo, data users.UsecaseInterface) {
	handler := &userDelivery{
		userUsecase: data,
	}

	e.GET("/users", handler.GetAllUser, middlewares.JWTMiddleware())
	e.GET("/users/me", handler.MyProfile, middlewares.JWTMiddleware())
	e.PUT("/users", handler.PutDataUser, middlewares.JWTMiddleware())
	e.DELETE("/manager/:id", handler.DeleteDataUser, middlewares.JWTMiddleware())
	e.POST("/manager", handler.PostDataUser) //<<=== Sementara tidak pakai token dulu

}

func (delivery *userDelivery) GetAllUser(c echo.Context) error {
	param := c.QueryParam("page")
	page, err := strconv.Atoi(param)
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("query param must be number"))
	}

	idToken, _ := middlewares.ExtractToken(c)

	data, errGet := delivery.userUsecase.GetAll(page, idToken)
	if errGet != nil {
		return c.JSON(400, helper.FailedResponseHelper("get all data failed"))
	} else if len(data) == 0 {
		return c.JSON(200, helper.SuccessResponseHelper("user data is empty"))
	}

	return c.JSON(200, helper.SuccessDataResponseHelper("success get data", toResponList(data)))
}

func (delivery *userDelivery) MyProfile(c echo.Context) error {

	idToken, _ := middlewares.ExtractToken(c)

	data, err := delivery.userUsecase.SelectMe(idToken)
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("failed get my profile"))
	}

	return c.JSON(200, helper.SuccessDataResponseHelper("get profile success", toRespon(data)))
}

func (delivery *userDelivery) PutDataUser(c echo.Context) error {

	idToken, _ := middlewares.ExtractToken(c)

	var updateRequest Request
	err := c.Bind(&updateRequest)
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("error bind"))
	}

	var updateData users.Core
	if updateRequest.Name != "" {
		updateData.Name = updateRequest.Name
	}
	if updateRequest.Email != "" {
		updateData.Email = updateRequest.Email
	}
	if updateRequest.Password != "" {
		updateData.Password = updateRequest.Password
	}
	if updateRequest.Role != "" {
		updateData.Role = updateRequest.Role
	}
	// if updateRequest.Team != "" {
	// 	updateData.Team = updateRequest.Team
	// }

	if idToken != 1 {
		if updateRequest.Role != "" || updateRequest.TeamID != 0 || updateRequest.ID != uint(idToken) {
			return c.JSON(400, helper.FailedResponseHelper("not have access"))
		}
	}
	updateData.ID = updateRequest.ID
	row := delivery.userUsecase.PutData(updateData)
	if row == -1 {
		return c.JSON(400, helper.FailedResponseHelper("failed to update data"))
	}

	return c.JSON(200, helper.SuccessResponseHelper("success update data"))

}

func (delivery *userDelivery) DeleteDataUser(c echo.Context) error {

	idToken, _ := middlewares.ExtractToken(c)
	if idToken != 1 {
		return c.JSON(400, helper.FailedResponseHelper("you not have access"))
	}

	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("param must be number"))
	}

	row := delivery.userUsecase.DeleteData(id)
	if row == -1 || row == 0 {
		return c.JSON(400, helper.FailedResponseHelper("delete data failed"))
	}

	return c.JSON(200, helper.SuccessResponseHelper("success delete data"))
}

func (delivery *userDelivery) PostDataUser(c echo.Context) error {

	var data Request
	err := c.Bind(&data)
	if err != nil || data.ID != 0 {
		return c.JSON(400, helper.FailedResponseHelper("error bindig data"))
	}

	// idToken, _ := middlewares.ExtractToken(c)
	// if idToken != 1 {
	// 	return c.JSON(400, helper.FailedResponseHelper("not have access"))
	// }

	row := delivery.userUsecase.PostData(data.toCoreReq())
	if row == -1 || row == 0 {
		return c.JSON(400, helper.FailedResponseHelper("failed insert data"))
	}

	return c.JSON(200, helper.SuccessResponseHelper("success insert data"))
}
