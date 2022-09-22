package delivery

import "immersiveProject/features/log/entity"

type LogRequest struct {
	Feedback  string `json:"feedback" form:"feedback"`
	StatusLog string `json:"statuslog" form:"statuslog"`
	File      string `json:"file" form:"file"`
}

func FromCoreRequest(logData LogRequest) entity.Log {
	return entity.Log{
		Feedback:  logData.Feedback,
		StatusLog: logData.StatusLog,
		File:      logData.File,
	}
}

func toCoreRequest(logData LogRequest) entity.Log {
	return entity.Log{
		Feedback:  logData.Feedback,
		StatusLog: logData.StatusLog,
		File:      logData.File,
	}
}
