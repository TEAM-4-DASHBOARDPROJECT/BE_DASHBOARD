package usecase

import "immersiveProject/features/class/entity"

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

func (usecase *classUsecase) Update(class entity.ClassEntity) (err error) {
	_, err = usecase.Repo.Update(class)

	if err != nil{
		return err
	}

	return nil
}

func (usecase *classUsecase) Delete(class entity.ClassEntity) (err error) {
	_, err = usecase.Repo.Delete(class)

	if err != nil{
		return err
	}

	return nil
}

func (usecase *classUsecase) GetClass() ([]entity.ClassEntity, error) {
	result, err := usecase.Repo.FindAll()
	
	return result, err
}