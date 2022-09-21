package data

import (
	"immersiveProject/features/class/entity"

	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	UserID     uint
	Name       string
	MulaiKelas string
	AkhirKelas string
	User       User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
}

func ModelToEntity(classModel Class) entity.ClassEntity {
	return entity.ClassEntity{
		ClassID:    classModel.ID,
		UserID:     classModel.UserID,
		Name:       classModel.Name,
		MulaiKelas: classModel.MulaiKelas,
		AkhirKelas: classModel.AkhirKelas,
	}
}

func EntityToModel(classEntity entity.ClassEntity) Class {
	return Class{
		UserID:     classEntity.UserID,
		Name:       classEntity.Name,
		MulaiKelas: classEntity.MulaiKelas,
		AkhirKelas: classEntity.AkhirKelas,
	}
}
