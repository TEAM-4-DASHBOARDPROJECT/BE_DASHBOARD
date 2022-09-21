package delivery

import (
	"immersiveProject/features/class/entity"
	"immersiveProject/middlewares"
	"immersiveProject/utils/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type classhandler struct{
	Usecase entity.UsecaseClass
}

func New(usecase entity.UsecaseClass) *classhandler {
	return &classhandler{
		Usecase: usecase,
	}
}

func (repo *classhandler) Create(c echo.Context) error {
	var userRequest ClassRequest

	id, err := middlewares.ExtractToken(c)
	if err != nil {
		return c.JSON(http.StatusForbidden, helper.FailedResponseHelper(err.Error()))
	}

	err = c.Bind(&userRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper(err.Error()))
	}
	userEntity := RequestToEntity(userRequest)
	userEntity.UserID = uint(id)

	err = repo.Usecase.Create(userEntity)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper(err.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponseHelper("Succses Create Class"))
}

func (repo *classhandler) Update(c echo.Context) error {
	var userRequest ClassRequest
	var ClassId int

	id, err := middlewares.ExtractToken(c)
	if err != nil {
		return c.JSON(http.StatusForbidden, helper.FailedResponseHelper(err.Error()))
	}

	ClassId, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper(err.Error()))
	}

	err = c.Bind(&userRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper(err.Error()))
	}

	userEntity := RequestToEntity(userRequest)
	userEntity.UserID = uint(id)
	userEntity.ClassID = uint(ClassId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper(err.Error()))
	}
	return c.JSON(http.StatusOK, helper.FailedResponseHelper("Succses update class"))
}

func (repo *classhandler) Delete(c echo.Context) error{
	userToken, errToken := middlewares.ExtractToken(c)
	if userToken == 0 || errToken != nil{
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("tolong tokennya tuan!"))
	}
	var classRemove ClassRequest
	classRemove.ClassID = userToken
	_, err := repo.Usecase.Delete(RequestToEntity(classRemove))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("gagal delete usernya"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponseHelper("delete berhasil"))
}


func (repo *classhandler) GetClass(c echo.Context) error{
	result, err := repo.Usecase.GetClass()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("failed get class"))
	}
	
	return c.JSON(http.StatusOK, helper.SuccessDataResponseHelper("succses get class", EntityList(result)))
}

// func (repo *classhandler) GetClass(c echo.Context) (err error){
// 	userToken, errToken := middlewares.ExtractToken(c)
// 	if userToken == 0 || errToken != nil {
// 		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("token invalid !"))
// 	}

// 	result, err := repo.Usecase.GetClass()
// 	classResult := ClassResponse{}

// 	for _, class := range result {
// 		if class.ClassID == uint(userToken) {
// 			classResult.ID = uint(userToken)
// 			classResult.Name = class.Name
// 			classResult.MulaiKelas = class.MulaiKelas
// 			classResult.AkhirKelas = class.AkhirKelas
// 		}
// 	}

// 	if err != nil || classResult.ID < 1{
// 		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("error get class"))
// 	}

// 	return c.JSON(http.StatusOK, helper.SuccessDataResponseHelper("succses", classResult))

// }