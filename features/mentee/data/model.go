package data

import (
	"immersiveProject/features/mentee"

	"gorm.io/gorm"
)

type Mentee struct {
	gorm.Model
	FullName          string
	Status            string
	Class             string
	Category          string
	Address           string
	HomeAddress       string
	Email             string `gorm:"unique"`
	Gender            string
	Telegram          string
	Phone             string
	EmergencyName     string
	EmergencyPhone    string
	EmergencyStatus   string
	EducationType     string
	EducationMajor    string
	EducationGraduate string
}

func fromCore(data mentee.Core) Mentee {
	dataModel := Mentee{
		FullName:          data.FullName,
		Status:            data.Status,
		Class:             data.Class,
		Category:          data.Category,
		Address:           data.Address,
		HomeAddress:       data.HomeAddress,
		Email:             data.Email,
		Gender:            data.Gender,
		Telegram:          data.Telegram,
		Phone:             data.Phone,
		EmergencyName:     data.EmergencyName,
		EmergencyPhone:    data.EmergencyPhone,
		EmergencyStatus:   data.EmergencyStatus,
		EducationType:     data.EducationType,
		EducationMajor:    data.EducationMajor,
		EducationGraduate: data.EducationGraduate,
	}
	return dataModel
}

func (data *Mentee) toCore() mentee.Core {
	return mentee.Core{
		FullName:          data.FullName,
		Status:            data.Status,
		Class:             data.Class,
		Category:          data.Category,
		Address:           data.Address,
		HomeAddress:       data.HomeAddress,
		Email:             data.Email,
		Gender:            data.Gender,
		Telegram:          data.Telegram,
		Phone:             data.Phone,
		EmergencyName:     data.EmergencyName,
		EmergencyPhone:    data.EmergencyPhone,
		EmergencyStatus:   data.EmergencyStatus,
		EducationType:     data.EducationType,
		EducationMajor:    data.EducationMajor,
		EducationGraduate: data.EducationGraduate,
	}
}
