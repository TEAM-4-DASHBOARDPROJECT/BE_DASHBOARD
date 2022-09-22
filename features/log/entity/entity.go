package entity

import "time"

type Log struct{
	LogID			int
	Feedback		string
	Status			string
	UrlFile			string
	CreatedAt		time.Time
	UpdatedAt		time.Time
}

type UsecaseLog	interface{
	InsertLog(logInsert Log)(row int, err error)
}

type InterfaceLog interface{
	CreateLog(logInsert Log)(row int, err error)
}