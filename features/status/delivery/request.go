package delivery

import "immersiveProject/features/status/entity"

type StatusRequest struct{
	StatusName	string	`json: "name" form: "name"`
}

func fromCore(data StatusRequest) entity.StatusEntity {
	return entity.StatusEntity{
		StatusName: data.StatusName,
	}
}