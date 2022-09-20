package delivery

import "immersiveProject/features/mentee"

type MenteeResponse struct {
	ID                uint   `json:"id"`
	FullName          string `json:"fullname"`
	StatusID          uint   `json:"status_id"`
	ClassID           uint   `json:"class_id"`
	Category          string `json:"category"`
	Address           string `json:"address,omitempty"`
	HomeAddress       string `json:"homeaddress,omitempty"`
	Email             string `json:"email,omitempty"`
	Gender            string `json:"gender"`
	Telegram          string `json:"telegram,omitempty"`
	Phone             string `json:"phone,omitempty"`
	EmergencyName     string `json:"emergencyname,omitempty"`
	EmergencyPhone    string `json:"emergencyphone,omitempty"`
	EmergencyStatus   string `json:"emergencystatus,omitempty"`
	EducationType     string `json:"educationtype,omitempty"`
	EducationMajor    string `json:"educationmajor,omitempty"`
	EducationGraduate string `json:"educationgraduate,omitempty"`
}

func fromCore(data mentee.Core) MenteeResponse {
	return MenteeResponse{
		ID:                data.ID,
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
