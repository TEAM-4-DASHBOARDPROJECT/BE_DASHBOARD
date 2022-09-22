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

	data := User{}
	tx := repo.db.Preload("Team").First(&data, token).Find(&data).Where("userid = ?", token)
	if tx.Error != nil {
		return users.Core{}, tx.Error
	}

	return data.toCore(), nil
}

func (repo *dataUser) SelectAll(page, token int) ([]users.Core, error) {

	limit := 5
	offset := ((page - 1) * limit)
	queryParam := repo.db.Limit(limit).Offset(offset).Preload("Team")

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

func (repo *dataUser) InsertData(input users.Core) (int, error) {
	UserModel := fromCore(input)
	tx := repo.db.Create(&UserModel)
	if tx.Error != nil {
		return 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("failed to insert data")
	}

	return int(tx.RowsAffected), nil
}
