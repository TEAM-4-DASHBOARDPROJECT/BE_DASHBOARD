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

func (repo *menteeData) SelectMentee(class string, status string, category string) ([]mentee.Core, error) {
	var dataMentee []Mentee
	//tx := repo.db.Where("class_id = ? OR status_id = ? OR category = ?", classID, statusID, category).Find(&dataMentee)

	if class == "" && status == "" && category != "" {
		var dataByCate []Mentee
		txCate := repo.db.Where("education_category = ?", category).Find(&dataByCate)
		if txCate.Error != nil {
			return nil, txCate.Error
		}
		return toCoreList(dataByCate), nil

	} else if class != "" && status == "" && category == "" {
		var dataByClass []Mentee
		txClass := repo.db.Where("name = ?", class).Preload("Classes").Find(&dataByClass)
		if txClass.Error != nil {
			return nil, txClass.Error
		}
		return toCoreList(dataMentee), nil

	} else if class == "" && status != "" && category == "" {
		var dataByStatus []Mentee
		txState := repo.db.Where("status = ?", status).Find(&dataByStatus)
		if txState.Error != nil {
			return nil, txState.Error
		}
		return toCoreList(dataByStatus), nil

	} else {
		txAll := repo.db.Find(&dataMentee)
		if txAll.Error != nil {
			return nil, txAll.Error
		}
		return toCoreList(dataMentee), nil
	}

}

func (repo *menteeData) DeleteData(id int) (int, error) {
	var deleteData Mentee
	tx := repo.db.Where("id = ?", id).Delete(&deleteData)
	if tx.Error != nil {
		return -1, tx.Error
	}
	return 1, nil
}
