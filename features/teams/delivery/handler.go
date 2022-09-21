package delivery

import (
	"immersiveProject/features/teams"
	"immersiveProject/middlewares"
	"immersiveProject/utils/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type teamHandler struct {
	Usecase teams.UsecaseTeam
}

func New(e *echo.Echo, usecase teams.UsecaseTeam) {
	handler := &teamHandler{
		Usecase: usecase,
	}

	e.POST("/teams", handler.Create, middlewares.JWTMiddleware())
	e.PUT("/teams/:id", handler.Update, middlewares.JWTMiddleware())
	e.GET("/teams", handler.GetTeam, middlewares.JWTMiddleware())
	e.DELETE("/teams", handler.Delete, middlewares.JWTMiddleware())

}

func (repo *teamHandler) Create(c echo.Context) error {
	var userRequest TeamRequest

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
	return c.JSON(http.StatusOK, helper.SuccessResponseHelper("Succses Create Teams"))
}

func (repo *teamHandler) Update(c echo.Context) error {
	var userRequest TeamRequest
	var teamId int

	id, err := middlewares.ExtractToken(c)
	if err != nil {
		return c.JSON(http.StatusForbidden, helper.FailedResponseHelper(err.Error()))
	}

	teamId, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper(err.Error()))
	}

	err = c.Bind(&userRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper(err.Error()))
	}

	userEntity := RequestToEntity(userRequest)
	userEntity.UserID = uint(id)
	userEntity.ID = uint(teamId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper(err.Error()))
	}
	return c.JSON(http.StatusOK, helper.FailedResponseHelper("Succses update team"))
}

func (repo *teamHandler) Delete(c echo.Context) error {
	var teamId int

	id, err := middlewares.ExtractToken(c)
	if err != nil {
		return c.JSON(http.StatusForbidden, helper.FailedResponseHelper(err.Error()))
	}

	teamId, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper(err.Error()))
	}

	userEntity := teams.Core{}
	userEntity.UserID = uint(id)
	userEntity.ID = uint(teamId)

	err = repo.Usecase.Delete(userEntity)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponseHelper("succses delete team"))
}

func (repo *teamHandler) GetTeam(c echo.Context) (err error) {
	result, err := repo.Usecase.GetTeam()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("failed get team"))
	}

	return c.JSON(http.StatusOK, helper.SuccessDataResponseHelper("succses get team", result))
}
