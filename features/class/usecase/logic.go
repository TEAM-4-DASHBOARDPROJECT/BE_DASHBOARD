package usecase

import (
	"errors"
	"immersiveProject/features/class/entity"
)

type classUsecase struct{
	Repo entity.RepoClass
}

func New(repo entity.RepoClass) *classUsecase {
	return &classUsecase{
		Repo: repo,
	}
}

func (usecase *classUsecase) Create(class entity.ClassEntity) (err error) {
	_, err = usecase.Repo.Insert(class)
	
	if err != nil {
		return err
	}

	return nil
}

func (usecase *classUsecase) Update(class entity.ClassEntity) (row int, err error) {
	classMap := make(map[string]interface{})
	if class.Name != ""{
		classMap["name"] = &class.Name
	}
	if class.JumlahKelas != ""{
		classMap["jumlah"] = &class.JumlahKelas
	}
	if class.MulaiKelas != ""{
		classMap["mulai"] = &class.MulaiKelas
	}
	if class.AkhirKelas != ""{
		classMap["akhir"] = &class.AkhirKelas
	}

	result, err := usecase.Repo.Update(class)
	if err != nil {
		return 0, errors.New("user already register? you can't update if your account not listed")
	}
	return result, err
}

func (usecase *classUsecase) Delete(class entity.ClassEntity) (row int, err error) {
	result, err := usecase.Repo.Delete(class)

	if err != nil{
		return -1, errors.New("what you delete? register account first and then try again delete account")
	}

	return result, err
}

func (usecase *classUsecase) GetClass() ([]entity.ClassEntity, error) {
	result, err := usecase.Repo.FindAll()
	
	return result, err
}