package delivery

import "immersiveProject/features/log/entity"

type LogRequest struct{
	Feedback	string	`json:"feedback" form:"feedback"`
	Status		string	`json:"status" form:"status"`
	UrlFile		string	`json:"urlfile" form:"urlfile"`
	UrlImage	string	`json:"urlimage" form:"urlimage"`
}

func FromCoreRequest(logData LogRequest) entity.Log {
	return entity.Log{
		Feedback: 	logData.Feedback,
		Status: 	logData.Status,
		UrlFile: 	logData.UrlFile,
		UrlImage:	logData.UrlImage,
	}
}

func toCoreRequest(logData LogRequest) entity.Log {
	return entity.Log{
		Feedback: 	logData.Feedback,
		Status: 	logData.Status,
		UrlFile: 	logData.UrlFile,
		UrlImage: 	logData.UrlImage,
	}
}