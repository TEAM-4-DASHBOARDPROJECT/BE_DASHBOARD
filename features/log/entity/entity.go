package entity

import "time"

type Log struct{
	LogID			int
	Feedback		string
	Status			string
	UrlFile			string
	UrlImage		string
	CreatedAt		time.Time
	UpdatedAt		time.Time
}

type Mentee struct{
	ID                uint
	FullName          string
	ClassID           uint
	Status            string
	Address           string
	HomeAddress       string
	Email             string
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

type UsecaseLog	interface{
	InsertLog(logInsert Log)(row int, err error)
	GetLog() (logInsert []Log, err error)
}

type InterfaceLog interface{
	FindLog() (logInsert []Log, err error)
	CreateLog(logInsert Log)(row int, err error)
}