package data

import (
	"immersiveProject/features/status/entity"

	"gorm.io/gorm"
)

type Status struct{
	gorm.Model
	Name	string
}

func FromEntity(data entity.StatusEntity) Status {
	statusModel := Status{
		Name: data.StatusName,
	}
	return statusModel
}

func (data *Status) ToEntity() entity.StatusEntity {
	return entity.StatusEntity{
		StatusID: 	int(data.ID),
		StatusName: data.Name,
	}
}

func CoreList(data []Status) []entity.StatusEntity {
	var StatusData []entity.StatusEntity
	for key := range data {
		StatusData = append(StatusData, data[key].ToEntity())
	}
	return StatusData
}