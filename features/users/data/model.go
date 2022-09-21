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
	Team     string
	Role     string
}

func (data *User) toCore() users.Core {
	return users.Core{
		ID:       data.ID,
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		Team:     data.Team,
		Role:     data.Role,
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
		Team:     data.Team,
		Role:     data.Role,
	}
}
