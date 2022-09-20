package data

import (
	"gorm.io/gorm"
	entity "immersiveProject/features/login/entity"
)

type User struct{
	gorm.Model
	Name		string
	Email		string
	Password	string
}

func ToCore(userModel User) entity.Login {
	return entity.Login{
		Id: userModel.ID,
		Email: userModel.Email,
		Password: userModel.Password,
	}
}