package delivery

import "immersiveProject/features/class/entity"

type ClassRequest struct {
	Name   string `json:"name" form:"name"`
	Jumlah string `json:"jumlah" form:"jumlah"`
	Mulai  string `json:"mulai" form:"mulai"`
	Akhir  string `json:"akhir" form:"akhir"`
}

func RequestToEntity(request ClassRequest) entity.ClassEntity {
	return entity.ClassEntity{
		Name:        request.Name,
		JumlahKelas: request.Jumlah,
		MulaiKelas:  request.Mulai,
		AkhirKelas:  request.Akhir,
	}
}
