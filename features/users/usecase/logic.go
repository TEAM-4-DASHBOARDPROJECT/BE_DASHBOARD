package usecase

import (
	"errors"
	"immersiveProject/features/users"

	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userData users.DataInterface
}

func New(data users.DataInterface) users.UsecaseInterface {
	return &userUsecase{
		userData: data,
	}
}

func (usecase *userUsecase) GetAll(page, token int) ([]users.Core, error) {

	data, err := usecase.userData.SelectAll(page, token)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (usecase *userUsecase) SelectMe(id int) (users.Core, error) {

	data, err := usecase.userData.GetMyProfile(id)
	if err != nil {
		return users.Core{}, err
	}

	return data, nil
}

func (usecase *userUsecase) PutData(data users.Core) int {

	if data.Password != "" {
		hash, _ := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
		data.Password = string(hash)
	}

	row := usecase.userData.UpdateData(data)
	if row == -1 {
		return -1
	}

	return row

}

func (usecase *userUsecase) DeleteData(id int) int {

	row := usecase.userData.DelData(id)
	if row == -1 {
		return -1
	}

	return row
}

func (usecase *userUsecase) PostData(data users.Core) (int, error) {

	if data.Password == "" || data.Email == "" || data.Name == "" || data.Role == "" || data.TeamID == 0 {
		return -1, errors.New("all data must be filled")
	}

	passWillBcrypt := []byte(data.Password)
	hash, err_hash := bcrypt.GenerateFromPassword(passWillBcrypt, bcrypt.DefaultCost)
	if err_hash != nil {
		return -2, errors.New("hashing password failed")
	}

	data.Password = string(hash)
	result, err := usecase.userData.InsertData(data)
	if err != nil {
		return 0, errors.New("failed to insert data")
	}

	return result, nil
}
