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
	//tx := repo.db.Where("class_id = ? OR status_id = ? OR category = ?", classID, statusID, category).Find(&dataMentee)

	if category != "" {
		var dataByCate []Results
		txCate := repo.db.Where("education_category = ?", category).Find(&dataByCate)
		if txCate.Error != nil {
			return nil, txCate.Error
		}
		return toCoreList(dataByCate), nil

	} else if status != "" {
		var dataByStatus []Results
		txState := repo.db.Where("status = ?", status).Find(&dataByStatus)
		if txState.Error != nil {
			return nil, txState.Error
		}
		return toCoreList(dataByStatus), nil

	} else if class != "" {
		var dataByClass []Results
		// txClass := repo.db.Where("name = ?", class).Preload("Mentees").Find(&Classes)
		txClass := repo.db.Model(&Class{}).Select("mentees.id, mentees.fullname, classes.name, mentees.status, mentees.address, mentees.homeaddress, mentees.email, mentees.gender, mentees.telegram, mentees.phone, mentees.emergencyname, mentees.emergencyphone, mentees.emergencystatus, mentees.educationcategory, mentees.educationmajor, mentees.educationgraduate, mentees.class_id").Joins("inner join mentees on mentees.class_id = classes.id").Where("mentees.class_name = ?", class).Scan(&dataByClass)
		if txClass.Error != nil {
			return nil, txClass.Error
		}
		return toCoreList(dataByClass), nil

	} else {
		var dataAll []Results
		txAll := repo.db.Find(&dataAll)
		if txAll.Error != nil {
			return nil, txAll.Error
		}
		return toCoreList(dataAll), nil
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
