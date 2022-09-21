package delivery

import (
	"immersiveProject/features/status/entity"
	"immersiveProject/utils/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type StatusHandler struct{
	statusUsecase entity.UsecaseStatus
}

func New(useCase entity.UsecaseStatus) *StatusHandler {
	return &StatusHandler{
		statusUsecase: useCase,
	}
}

func (handler *StatusHandler) GetStatus(c echo.Context) error{
	result, err := handler.statusUsecase.GetStatus()
	if err != nil{
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("internal server error"))
	}
	return c.JSON(http.StatusOK, helper.SuccessDataResponseHelper("succses get status", CoreList(result)))
}