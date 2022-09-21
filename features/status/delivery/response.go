package delivery

import "immersiveProject/features/status/entity"

type StatusResponse struct{
	StatusID	uint	`json: "id"`
	StatusName	string	`json: "name"`
}

func FromCore(data entity.StatusEntity) StatusResponse {
	return StatusResponse{
		StatusID:  uint(data.StatusID),
		StatusName: data.StatusName,
	}
}

func CoreList(data []entity.StatusEntity) []StatusResponse {
	var Response []StatusResponse
	for _, v := range data {
		Response = append(Response, FromCore(v))
	}
	return Response
}

func CoreResponse(data entity.StatusEntity) StatusResponse {
	Response := StatusResponse{
		StatusID: uint(data.StatusID),
		StatusName: data.StatusName,
	}
	return Response
}