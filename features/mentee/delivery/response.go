package delivery

import "immersiveProject/features/mentee"

type MenteeResponse struct {
	ID                int    `json:"id"`
	FullName          string `json:"fullname"`
	Status            string `json:"status"`
	Class             string `json:"class"`
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
