package delivery

import (
	"immersiveProject/features/teams"
)

type TeamRequest struct {
	Name string `json:"name" form:"name"`
}

func RequestToEntity(request TeamRequest) teams.Core {
	return teams.Core{
		Name: request.Name,
	}
}
