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

type UsecaseLog	interface{
	InsertLog(logInsert Log)(row int, err error)
	GetLog() (logInsert []Log, err error)
}

type InterfaceLog interface{
	FindLog() (logInsert []Log, err error)
	CreateLog(logInsert Log)(row int, err error)
}