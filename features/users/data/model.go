package data

import (
	"immersiveProject/features/users"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique"`
	Password string
	TeamID   uint
	Role     string
	Status   string
	Team     Team `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Team struct {
	gorm.Model
	Name string
	User []User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (data *User) toCore() users.Core {
	return users.Core{
		ID:       data.ID,
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		Role:     data.Role,
		Status:   data.Status,
		Team: users.Team{
			ID:   int(data.TeamID),
			Name: data.Team.Name,
		},
	}
}

func toCoreList(data []User) []users.Core {
	var dataAll []users.Core
	for key := range data {
		dataAll = append(dataAll, data[key].toCore())
	}

	return dataAll
}

func fromCore(data users.Core) User {
	return User{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		Role:     data.Role,
		Status:   data.Status,
		TeamID:   uint(data.Team.ID),
	}
}
