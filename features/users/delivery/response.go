package delivery

import "immersiveProject/features/users"

type Response struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Team   string `json:"team"`
	Role   string `json:"role"`
	Status string `json:"status"`
}

func toRespon(data users.Core) Response {
	return Response{
		ID:     data.ID,
		Name:   data.Name,
		Email:  data.Email,
		Team:   data.Team,
		Role:   data.Role,
		Status: data.Status,
	}
}

func toResponList(data []users.Core) []Response {
	var dataAll []Response
	for key := range data {
		dataAll = append(dataAll, toRespon(data[key]))
	}

	return dataAll
}
