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
	e.PUT("/mentee/:id", handler.PostMentee, middlewares.JWTMiddleware())
	e.GET("/mentee", handler.PostMentee, middlewares.JWTMiddleware())
	e.DELETE("/mentee/:id", handler.PostMentee, middlewares.JWTMiddleware())

}

func (delivery *MenteeDelivery) PostMentee(c echo.Context) error {
	var dataRegister MenteeRequest
	errBind := c.Bind(&dataRegister)
	if errBind != nil {
		return c.JSON(400, helper.FailedResponseHelper("error bind"))
	}

	// if dataRegister.Gender != "male" || dataRegister.Gender != "female" {
	// 	return c.JSON(400, helper.FailedResponseHelper("gender must be male or female"))
	// }

	// if dataRegister.EmergencyStatus != "orang tua" || dataRegister.EmergencyStatus != "saudara kandung" {
	// 	return c.JSON(400, helper.FailedResponseHelper("emergency status must be orang tua or saudara kandung"))
	// }

	// if dataRegister.EducationCategory != "informatics" || dataRegister.EducationCategory != "non-informatics" {
	// 	return c.JSON(400, helper.FailedResponseHelper("education category must be informatics or non-informatics"))
	// }

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

func (delivery *MenteeDelivery) GetMentee(c echo.Context) error {
	var dataMentee []mentee.Core

	query := c.QueryParam("class")
	query2 := c.QueryParam("status")
	query3 := c.QueryParam("category")

	// idClass, err := strconv.Atoi(query)
	// if err != nil {
	// 	return c.JSON(400, helper.FailedResponseHelper("query param must be number"))
	// }
	// classID := idClass

	// idStatus, err := strconv.Atoi(query2)
	// if err != nil {
	// 	return c.JSON(400, helper.FailedResponseHelper("query param must be number"))
	// }
	// statusID := idStatus

	data, errGet := delivery.menteeUsecase.GetMentee(query, query2, query3)
	if errGet != nil {
		return c.JSON(400, helper.FailedResponseHelper("error get all data"))
	} else if len(data) == 0 {
		return c.JSON(400, helper.FailedResponseHelper("data still empty"))
	}

	return c.JSON(200, helper.SuccessDataResponseHelper("success get data", fromCoreList(dataMentee)))
}

func (delivery *MenteeDelivery) DeleteMentee(c echo.Context) error {
	id := c.Param("id")
	idCnv, errId := strconv.Atoi(id)
	if errId != nil {
		return c.JSON(400, helper.FailedResponseHelper("param must be number"))
	}

	row, err := delivery.menteeUsecase.DeleteMentee(idCnv)
	if err != nil || row != 1 {
		return c.JSON(400, helper.FailedResponseHelper("failed delete"))
	}
	return c.JSON(200, helper.SuccessResponseHelper("success delete"))
}
