package data

import (
	"immersiveProject/features/teams"

	"gorm.io/gorm"
)

type Team struct {
	gorm.Model
	UserID uint
	Name   string
}

func ModelToEntity(data Team) teams.Core {
	return teams.Core{
		ID:   data.ID,
		Name: data.Name,
	}
}

func entityToModel(core teams.Core) Team {
	return Team{
		UserID: core.UserID,
		Name:   core.Name,
	}
}
