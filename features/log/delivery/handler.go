package delivery

import (
	// "immersiveProject/config"
	"immersiveProject/features/log/entity"
	"log"

	// "immersiveProject/features/users/data"
	"immersiveProject/middlewares"
	"immersiveProject/utils/helper"
	"net/http"

	// "strconv"
	// "time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/labstack/echo/v4"
)

type loghandler struct {
	LogInterface entity.InterfaceLog
	conn		*session.Session
}

func New(log entity.InterfaceLog) *loghandler {
	aws := &session.Session{}
	return &loghandler{
		LogInterface: log,
		conn: 	aws,
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
	logToken, errToken := middlewares.ExtractToken(c)
	if logToken == 0 || errToken != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("failed insert data"))
	}
	var logs LogResponse
	bind := c.Bind(&logs)

	if bind != nil {
		log.Print("fail bind")
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("internal server error"))
	}

	file, err := c.FormFile("file")
	if err != nil {
		log.Println(err)
	}
	link := helper.DoUpload(handler.conn, *file, file.Filename)
	logs.File = link

	if err != nil{
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("internal server error"))
	}
	return c.JSON(http.StatusOK, helper.SuccessDataResponseHelper("Succses", link))
	
}



// var response LogResponse
// 	var fileType, fileName string
// 	var fileSize int64
// 	isSuccses := true
// 	file, err := c.FormFile("file")

// 	if err != nil {
// 		isSuccses = false
// 	} else{
// 		src, err := file.Open()
// 		if err != nil {
// 			isSuccses = false
// 		} else {
// 			fileByte, _ := ioutil.ReadAll(src)
// 			fileType = http.DetectContentType(fileByte)

// 			if fileType == "application/pdf"{
// 				fileName = strconv.FormatInt(time.Now().Unix(), 10)+".pdf"
// 			} else if fileType == "jpg"{
// 				fileName = strconv.FormatInt(time.Now().Unix(), 10)+".jpg"
// 			} else {
// 				fileName = strconv.FormatInt(time.Now().Unix(), 10)+".jpeg"
// 			}

// 			err = ioutil.WriteFile(fileName, fileByte, 0777)
// 			if err != nil{
// 				isSuccses = false
// 			} else {
// 				fileSize = file.Size
// 			}
// 		}
// 	}

// 	if isSuccses {
// 		return c.JSON(http.StatusOK, helper.SuccessDataResponseHelper("Succses upload file", file))
// 	} else {
// 		response = LogResponse{
// 			Message: "Failed upload file",
// 		}
// 	}


// logToken, errToken := middlewares.ExtractToken(c)
	// if logToken == 0 || errToken != nil {
	// 	return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("failed insert data"))
	// }

	// logs := LogRequest{}
	// errBind := c.Bind(&logs)
	// if errBind != nil {
	// 	return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("error bind log"))
	// }

	// fileData, fileInfo, fileErr := c.Request().FormFile("file")
	// if fileErr == http.ErrMissingFile || fileErr != nil {
	// 	return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("failed to get file"))
	// }

	// fileExtension, errFileExtention := helper.CheckfileExtension(fileInfo.Filename, config.ContentDocuments)
	// if errFileExtention != nil {
	// 	return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("file extension error"))
	// }

	// errFileSize := helper.CheckFileSize(fileInfo.Size, config.ContentDocuments)
	// if errFileSize != nil {
	// 	return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("file size error"))
	// }

	// filename := strconv.Itoa(logToken) + "" + logs.File + time.Now().Format("2006-01-02 15:04:05") + fileExtension
	// file, errUploadFile := helper.UploadPDFToS3(config.ContentDocuments, filename, config.ContentDocuments, fileData)

	// if errUploadFile != nil {
	// 	return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("failed to upload file"))
	// }

	// logsCore := toCoreRequest(logs)
	// logsCore.LogID = logToken
	// logsCore.File = file

	// _, err := handler.LogInterface.CreateLog(logsCore)
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("failed insert logs"))
	// }
	// return c.JSON(http.StatusOK, helper.SuccessResponseHelper("insert logs succses"))