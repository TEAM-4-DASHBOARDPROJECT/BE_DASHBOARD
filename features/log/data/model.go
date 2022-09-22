package data

import (
	"immersiveProject/features/log/entity"

	"gorm.io/gorm"
)

type Log struct {
	gorm.Model
	UserID   uint
	MenteeID uint
	Feedback string
	Status   string
	UrlFile  string
	UrlImage string
	User     User
	Mentee   Mentee
}

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique"`
	Password string
	Team     string
	Role     string
	Status   string
}

type Mentee struct {
	gorm.Model
	UserID            uint
	MenteeID          uint
	Status            string
	Address           string
	HomeAddress       string
	Email             string `gorm:"unique"`
	Gender            string
	Telegram          string
	Phone             string
	EmergencyName     string
	EmergencyPhone    string
	EmergencyStatus   string
	EducationCategory string
	EducationMajor    string
	EducationGraduate string
}

func FromCore(logCore entity.Log) Log {
	logModel := Log{
		Feedback: logCore.Feedback,
		Status:   logCore.Status,
		UrlFile:  logCore.UrlFile,
		UrlImage: logCore.UrlImage,
	}
	return logModel
}

func (logCore *Log) ToCore() entity.Log {
	return entity.Log{
		LogID:    int(logCore.ID),
		Feedback: logCore.Feedback,
		Status:   logCore.Status,
		UrlFile:  logCore.UrlFile,
		UrlImage: logCore.UrlImage,
	}
}

func CoreList(logCore []Log) []entity.Log {
	var LogCore []entity.Log
	for key := range logCore {
		LogCore = append(LogCore, logCore[key].ToCore())
	}
	return LogCore
}
