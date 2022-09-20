package usecase

import (
	"immersiveProject/features/mentee"
)

type menteeUsecase struct {
	menteeData mentee.DataInterface
}

func New(data mentee.DataInterface) mentee.UsecaseInterface {
	return &menteeUsecase{
		menteeData: data,
	}
}

func (usecase *menteeUsecase) PostMentee(data mentee.Core) (int, error) {
	row, err := usecase.menteeData.AddMentee(data)
	if err != nil {
		return -1, err
	}
	return row, nil
}
