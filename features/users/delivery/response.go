package delivery

import "immersiveProject/features/users"

type Response struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	// TeamID     int `json:"team"`
	Role   string `json:"role"`
	Status string `json:"status"`
	Team   Team
}

type Team struct {
	ID   int
	Name string
}

func toRespon(data users.Core) Response {
	return Response{
		ID:       data.ID,
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		// Team:     data.Team,
		Role:   data.Role,
		Status: data.Status,
		Team: Team{
			ID:   data.Team.ID,
			Name: data.Team.Name,
		},
	}
}

func toResponList(data []users.Core) []Response {
	var dataAll []Response
	for key := range data {
		dataAll = append(dataAll, toRespon(data[key]))
	}

	return dataAll
}