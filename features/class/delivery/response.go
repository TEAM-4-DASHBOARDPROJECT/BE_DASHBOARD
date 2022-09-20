package delivery

import "immersiveProject/features/class/entity"

type ClassResponse struct{
	ID			uint
	Name		string
	MulaiKelas	string
	AkhirKelas	string
}

func EntityToClassResponse(classEntity entity.ClassEntity) ClassResponse {
	return ClassResponse{
		ID: 		classEntity.ClassID,
		Name: 		classEntity.Name,
		MulaiKelas: classEntity.MulaiKelas,
		AkhirKelas: classEntity.AkhirKelas,
	}
}