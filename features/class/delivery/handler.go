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

	if err := c.Validate(userRequest); err != nil {
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

	if err := c.Validate(userRequest); err != nil{
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