package delivery

import "immersiveProject/features/mentee"

type MenteeRequest struct {
	FullName          string `json:"name" form:"name"`
	StatusID          uint   `json:"status_id" form:"status_id"`
	ClassID           uint   `json:"class_id" form:"class_id"`
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
		StatusID:          data.StatusID,
		ClassID:           data.ClassID,
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
