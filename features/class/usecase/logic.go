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

func (usecase *classUsecase) GetClass(class entity.ClassEntity) (result []entity.ClassEntity, err error) {
	result, err = usecase.Repo.FindAll(class)
	
	return result, err
}

func (usecase *classUsecase) GetSingleClass(class entity.ClassEntity)(result entity.ClassEntity, err error) {
	result, err = usecase.Repo.Find(class)

	return result, err
}