package delivery

import (
	// "immersiveProject/config"
	"immersiveProject/features/log/entity"
	// "immersiveProject/features/mentee/delivery"
	"immersiveProject/utils/helper"
	"net/http"
	// "strconv"
	// "time"

	"github.com/labstack/echo/v4"
)

type loghandler struct {
	LogInterface entity.InterfaceLog
}

func New(log entity.InterfaceLog) *loghandler {
	return &loghandler{
		LogInterface: log,
	}
}

func (handler *loghandler) FindLog(c echo.Context) error {
	result, err := handler.LogInterface.FindLog()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("internal server error"))
	}
	return c.JSON(http.StatusOK, helper.SuccessDataResponseHelper("succses get Log", CoreList(result)))
}

func (handler *loghandler) Createlog(c echo.Context) error {
	// var response 
	return nil
}
