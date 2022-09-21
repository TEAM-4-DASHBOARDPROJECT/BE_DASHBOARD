package data

import (
	"immersiveProject/features/teams"

	"gorm.io/gorm"
)

type Team struct {
	gorm.Model
	UserID uint
	Name   string
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	// Team     Team
}

func ModelToEntity(data Team) teams.Core {
	return teams.Core{
		ID:     data.ID,
		Name:   data.Name,
		UserID: data.User.ID,
	}
}

func entityToModel(core teams.Core) Team {
	return Team{
		UserID: core.UserID,
		Name:   core.Name,
	}
}

// func (data *Team) toCore() teams.Core {
// 	return teams.Core{
// 		ID:   int(data.ID),
// 		Name: data.Name,
// 		User: teams.User{
// 			ID:   int(data.UserID),
// 			Name: data.User.Name,
// 		},
// 	}
// }

// func toCoreList(data []Team) []teams.Core {
// 	result := []teams.Core{}
// 	for key := range data {
// 		result = append(result, data[key].toCore())
// 	}
// 	return result
// }

// func FromCore(core teams.Core) Team {
// 	return Team{
// 		Name:   core.Name,
// 		UserID: int(core.User.ID),
// 	}
// }
