package delivery

import (
	"immersiveProject/features/teams"
)

type TeamResponse struct {
	ID   uint
	Name string
}

func EntityToClassResponse(data teams.Core) TeamResponse {
	return TeamResponse{
		ID:   data.ID,
		Name: data.Name,
	}
}
