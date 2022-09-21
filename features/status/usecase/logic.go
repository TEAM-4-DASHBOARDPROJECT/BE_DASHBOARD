package usecase

import "immersiveProject/features/status/entity"

type statusUsecase struct{
	statusData entity.RepoStatus
}

func New(logic entity.RepoStatus) entity.RepoStatus {
	return &statusUsecase{
		statusData: logic,
	}
}

func (status *statusUsecase)SelectStatus() ([]entity.StatusEntity, error) {
	result, err := status.statusData.SelectStatus()
	return result, err
}