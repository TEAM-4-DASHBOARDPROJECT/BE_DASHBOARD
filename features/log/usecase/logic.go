package usecase

import (
	"errors"
	"immersiveProject/features/log/entity"
)

type logLogic struct{
	logData 	entity.UsecaseLog
}

func New(logic entity.UsecaseLog) entity.UsecaseLog {
	return &logLogic{
		logData: logic,
	}
}

func (logic *logLogic) InsertLog(logCreate entity.Log) (int, error){
	if logCreate.Feedback == "" || logCreate.Status == "" || logCreate.UrlFile == "" {
		return -1, errors.New("log not fuel field")
	}
	
	result, err := logic.logData.InsertLog(logCreate)
	if err != nil{
		return 0, errors.New("failed to insert data")
	}
	return result, nil
}