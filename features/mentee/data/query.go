package data

import (
	"immersiveProject/features/mentee"

	"gorm.io/gorm"
)

type menteeData struct {
	db *gorm.DB
}

func New(db *gorm.DB) mentee.DataInterface {
	return &menteeData{
		db: db,
	}
}

func (repo *menteeData) AddMentee(data mentee.Core) (int, error) {
	dataModel := fromCore(data)
	tx := repo.db.Create(&dataModel)
	if tx.Error != nil {
		return -1, tx.Error
	}
	return int(tx.RowsAffected), nil
}

func (repo *menteeData) UpdateMentee(id int, newData mentee.Core) (int, error) {
	dataModel := fromCore(newData)
	tx := repo.db.Model(&Mentee{}).Where("id = ?", id).Updates(dataModel)
	if tx.Error != nil {
		return -1, tx.Error
	}
	return 1, nil
}

func (repo *menteeData) SelectMentee(get string) ([]mentee.Core, error) {
	var dataMentee []Mentee
	tx := repo.db.Where("name = ?", get).Find(&dataMentee)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return toCoreList(dataMentee), nil
}
