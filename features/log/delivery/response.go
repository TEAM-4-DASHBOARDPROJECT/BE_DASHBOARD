package delivery

import "immersiveProject/features/log/entity"

type LogResponse struct{
	Feedback	string `json:"feedback"`
	Status		string `json:"status"`
	UrlFile		string `json:"urlfile"`
	UrlImage	string `json:"urlimage"`
}

func FromCore(logData entity.Log) LogResponse {
	return LogResponse{
		Feedback: 	logData.Feedback,
		Status: 	logData.Status,
		UrlFile: 	logData.UrlFile,
		UrlImage: 	logData.UrlImage,
	}
}

func CoreList(logData []entity.Log) []LogResponse {
	var ResponseLog []LogResponse
	for _, v := range logData {
		ResponseLog = append(ResponseLog, FromCore(v))
	}
	return ResponseLog
}

func CoreResponse(logData entity.Log) LogResponse {
	ResponseLog := LogResponse{
		Feedback: logData.Feedback,
		Status: logData.Status,
		UrlFile: logData.UrlFile,
		UrlImage: logData.UrlImage,
	}
	return ResponseLog
}