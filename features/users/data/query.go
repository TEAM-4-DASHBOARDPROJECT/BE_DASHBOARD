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

func (repo *dataUser) SelectAll(page, token int) ([]users.Core, error) {

	limit := 5
	offset := ((page - 1) * limit)
	queryBuider := repo.db.Limit(limit).Offset(offset)

	var data []User
	tx := queryBuider.Find(&data).Order("name ASC")
	if tx.Error != nil {
		return nil, tx.Error
	}

	for _, v := range data {
		if v.ID == uint(token) {
			return toCoreList(data), nil
		}
	}

	return nil, errors.New("not have access")
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
