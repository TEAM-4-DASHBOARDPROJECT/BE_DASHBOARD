package delivery

import "immersiveProject/features/class/entity"

type ClassRequest struct{
	Name		string	`json:"name" form:"name"`
	MulaiKelas	string	`json:"mulai" form:"mulai"`
	AkhirKelas	string	`json:"akhir" form:"akhir"`
}

func RequestToEntity(request ClassRequest) entity.ClassEntity {
	return entity.ClassEntity{
		Name: 		request.Name,
		MulaiKelas: request.MulaiKelas,
		AkhirKelas: request.AkhirKelas,
	}
}