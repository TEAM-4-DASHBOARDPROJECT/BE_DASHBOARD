package repository

import (
	datas "immersiveProject/features/status/data"
	status "immersiveProject/features/status/entity"

	"gorm.io/gorm"
)

type RepoStatus struct{
	db *gorm.DB
}

func new(db *gorm.DB) status.RepoStatus {
	return &RepoStatus{
		db: db,
	}
}

func (repo *RepoStatus) SelectStatus()([]status.StatusEntity, error) {
	var statusRepo []datas.Status
	tx := repo.db.Find(&statusRepo)
	if tx.Error != nil{
		return nil, tx.Error
	}
	statused := datas.CoreList(statusRepo)
	return statused, nil
}