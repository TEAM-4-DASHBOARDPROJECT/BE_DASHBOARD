package delivery

import (
	"immersiveProject/config"
	"immersiveProject/features/log/entity"
	"immersiveProject/middlewares"
	"immersiveProject/utils/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type loghandler struct{
	LogInterface	entity.InterfaceLog
}

func New(log entity.InterfaceLog) *loghandler {
	return &loghandler{
		LogInterface: log,
	}
}

func (handler *loghandler) Createlog(c echo.Context) error {
	logToken, errToken := middlewares.ExtractToken(c)
	if logToken == 0 || errToken != nil{
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("token invalid !"))
	}

	logs := LogRequest{}
	errBind := c.Bind(&logs)
	if errBind != nil{
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("error bind log"))
	}

	imageData, imageInfo, imageErr := c.Request().FormFile("urlimage")
	if imageErr == http.ErrMissingFile || imageErr != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("failed to get image"))
	}

	imageExtension, errImageExtension := helper.CheckfileExtension(imageInfo.Filename, config.ContentImage)
	if errImageExtension != nil{
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("image extension error"))
	}

	errImageSize := helper.CheckFileSize(imageInfo.Size, config.ContentImage)
	if errImageSize != nil{
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("image size error"))
	}

	imageName := strconv.Itoa(logToken) + " " + "Feedback" + imageExtension

	image, errUploadImage := helper.UploadFileToS3(config.LogsImages, imageName, config.ContentImage, imageData)

	if errUploadImage != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("failed to upload file"))
	}

	fileData, fileInfo, fileErr := c.Request().FormFile("urlfile")
	if fileErr == http.ErrMissingFile || fileErr != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("failed to get file"))
	}

	fileExtension, errFileExtention := helper.CheckfileExtension(fileInfo.Filename, config.ContentDocuments)
	if errFileExtention != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("file extension error"))
	}

	errFileSize := helper.CheckFileSize(fileInfo.Size, config.ContentDocuments)
	if errFileSize != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("file size error"))
	}

	filename := strconv.Itoa(logToken) + "_ " + "Feedback" + fileExtension
	file, errUploadFile := helper.UploadPDFToS3(config.ContentDocuments, filename, config.ContentDocuments, fileData)

	if errUploadFile != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("failed to upload file"))
	}

	
	logsCore := toCoreRequest(logs)
	logsCore.LogID = logToken
	logsCore.UrlFile = file
	logsCore.UrlImage = image

	_, err := handler.LogInterface.CreateLog(logsCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("failed insert logs"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponseHelper("insert logs succses"))
}