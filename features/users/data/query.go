package data

import (
	"errors"
	"immersiveProject/features/users"

	"gorm.io/gorm"
)

type dataUser struct {
	db *gorm.DB
}

func New(db *gorm.DB) users.DataInterface {
	return &dataUser{
		db: db,
	}
}
func (repo *dataUser) GetMyProfile(token int) (users.Core, error) {

	var data User
	tx := repo.db.First(&data, token)
	if tx.Error != nil {
		return users.Core{}, tx.Error
	}

	return data.toCore(), nil
}

func (repo *dataUser) SelectAll(page, token int) ([]users.Core, error) {

	limit := 5
	offset := ((page - 1) * limit)
	queryParam := repo.db.Limit(limit).Offset(offset)

	var data []User

	if page > 0 {
		txA := queryParam.Find(&data).Order("name ASC")
		if txA.Error != nil {
			return nil, txA.Error
		}
	} else {
		txB := repo.db.Find(&data).Order("name ASC")
		if txB.Error != nil {
			return nil, txB.Error
		}
	}

	for _, v := range data {
		if v.ID == uint(token) {
			return toCoreList(data), nil
		}
	}

	return nil, errors.New("you not have access")
}

func (repo *dataUser) UpdateData(data users.Core) int {

	newData := fromCore(data)

	tx := repo.db.Model(&User{}).Where("id = ? ", int(data.ID)).Updates(&newData)
	if tx.Error != nil {
		return -1
	}

	return int(tx.RowsAffected)

}

func (repo *dataUser) DelData(id int) int {

	tx := repo.db.Unscoped().Where("id = ? ", id).Delete(&User{})
	if tx.Error != nil {
		return -1
	}

	return int(tx.RowsAffected)
}

func (repo *dataUser) InsertData(data users.Core) int {

	var dataInsert = fromCore(data)
	tx := repo.db.Create(&dataInsert)
	if tx.Error != nil {
		return -1
	}

	return int(tx.RowsAffected)
}
