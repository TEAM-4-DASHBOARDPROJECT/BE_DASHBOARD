package delivery

import (
	"immersiveProject/features/mentee"
	"immersiveProject/middlewares"
	"immersiveProject/utils/helper"
	"strconv"

	"github.com/labstack/echo/v4"
)

type MenteeDelivery struct {
	menteeUsecase mentee.UsecaseInterface
}

func New(e *echo.Echo, usecase mentee.UsecaseInterface) {
	handler := &MenteeDelivery{
		menteeUsecase: usecase,
	}

	e.POST("/mentee", handler.PostMentee, middlewares.JWTMiddleware())
	e.PUT("/mentee/{:id}", handler.PostMentee, middlewares.JWTMiddleware())

}

func (delivery *MenteeDelivery) PostMentee(c echo.Context) error {
	var dataRegister MenteeRequest
	errBind := c.Bind(&dataRegister)
	if errBind != nil {
		return c.JSON(400, helper.FailedResponseHelper("error bind"))
	}

	if dataRegister.Gender != "male" || dataRegister.Gender != "female" {
		return c.JSON(400, helper.FailedResponseHelper("gender must be male or female"))
	}

	if dataRegister.EmergencyStatus != "orang tua" || dataRegister.EmergencyStatus != "saudara kandung" {
		return c.JSON(400, helper.FailedResponseHelper("gender must be orang tua or saudara kandung"))
	}

	if dataRegister.EducationType != "informatics" || dataRegister.EducationType != "non-informatics" {
		return c.JSON(400, helper.FailedResponseHelper("gender must be informatics or non-informatics"))
	}

	row, err := delivery.menteeUsecase.PostMentee(toCore(dataRegister))
	if err != nil {
		return c.JSON(500, helper.FailedResponseHelper("error insert data"))
	}
	if row != 1 {
		return c.JSON(500, helper.FailedResponseHelper("error insert data"))
	}

	return c.JSON(201, helper.SuccessResponseHelper("success insert data"))

}

func (delivery *MenteeDelivery) PutMentee(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("param must be number"))
	}

	var dataUpdate MenteeRequest
	errBind := c.Bind(&dataUpdate)
	if errBind != nil {
		return c.JSON(400, helper.FailedResponseHelper("error bind data"))
	}

	row, err := delivery.menteeUsecase.PutMentee(id, toCore(dataUpdate))
	if err != nil {
		return c.JSON(500, helper.FailedResponseHelper("error update data"))
	}
	if row != 1 {
		return c.JSON(500, helper.FailedResponseHelper("error update data"))
	}

	return c.JSON(201, helper.SuccessResponseHelper("success update data"))
}
