package delivery

import "immersiveProject/features/users"

type Request struct {
	ID       uint   `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	TeamID   int    `json:"team" form:"team"`
	Role     string `json:"role" form:"role"`
	Status   string `json:"status" form:"status"`
}

func (req *Request) toCoreReq() users.Core {
	return users.Core{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		// TeamID:   req.TeamID,
		Role:   req.Role,
		Status: req.Status,
		Team: users.Team{
			ID: req.TeamID,
		},
	}
}
