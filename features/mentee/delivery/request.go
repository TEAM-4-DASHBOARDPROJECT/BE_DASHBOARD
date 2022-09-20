package delivery

import "immersiveProject/features/mentee"

type MenteeRequest struct {
	FullName          string `json:"name" form:"name"`
	Status            string `json:"status" form:"status"`
	Class             string `json:"class" form:"class"`
	Category          string `json:"category" form:"category"`
	Address           string `json:"address" form:"address"`
	HomeAddress       string `json:"homeaddress" form:"homeaddress"`
	Email             string `json:"email" form:"email"`
	Gender            string `json:"gender" form:"gender"`
	Telegram          string `json:"telegram" form:"telegram"`
	Phone             string `json:"phone" form:"phone"`
	EmergencyName     string `json:"emergencyname" form:"emergencyname"`
	EmergencyPhone    string `json:"emergencyphone" form:"emergencyphone"`
	EmergencyStatus   string `json:"emergencystatus" form:"emergencystatus"`
	EducationType     string `json:"educationtype" form:"educationtype"`
	EducationMajor    string `json:"educationmajor" form:"educationmajor"`
	EducationGraduate string `json:"educationgraduate" form:"educationgraduate"`
}

func toCore(data MenteeRequest) mentee.Core {
	return mentee.Core{
		FullName:          data.FullName,
		Status:            data.FullName,
		Class:             data.FullName,
		Category:          data.FullName,
		Address:           data.FullName,
		HomeAddress:       data.FullName,
		Email:             data.FullName,
		Gender:            data.FullName,
		Telegram:          data.FullName,
		Phone:             data.FullName,
		EmergencyName:     data.FullName,
		EmergencyPhone:    data.FullName,
		EmergencyStatus:   data.FullName,
		EducationType:     data.FullName,
		EducationMajor:    data.FullName,
		EducationGraduate: data.FullName,
	}
}
