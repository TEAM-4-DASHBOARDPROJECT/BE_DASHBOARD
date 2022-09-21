package usecase

import (
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

func (usecase *userUsecase) PostData(data users.Core) int {

	if data.Password == "" || data.Email == "" || data.Name == "" || data.Role == "" || data.Team == "" {
		return -1
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	data.Password = string(hash)

	row := usecase.userData.InsertData(data)
	if row == -1 {
		return -1
	}

	return row
}
