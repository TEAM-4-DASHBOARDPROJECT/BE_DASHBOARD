package usecase

import (
	"immersiveProject/features/teams"
)

type teamsUsecase struct {
	Repo teams.RepoTeam
}

func New(repo teams.RepoTeam) *teamsUsecase {
	return &teamsUsecase{
		Repo: repo,
	}
}

func (usecase *teamsUsecase) Create(data teams.Core) (err error) {
	_, err = usecase.Repo.Insert(data)

	if err != nil {
		return err
	}

	return nil
}

func (usecase *teamsUsecase) Update(data teams.Core) (err error) {
	_, err = usecase.Repo.Update(data)

	if err != nil {
		return err
	}

	return nil
}

func (usecase *teamsUsecase) Delete(data teams.Core) (err error) {
	_, err = usecase.Repo.Delete(data)

	if err != nil {
		return err
	}

	return nil
}

func (usecase *teamsUsecase) GetTeam() (result []teams.Core, err error) {
	result, err = usecase.Repo.FindAll()

	return result, err
}

// type TeamUsecase struct {
// 	Repo teams.UsecaseTeam
// }

// func New(repo teams.UsecaseTeam) *TeamUsecase {
// 	return &TeamUsecase{
// 		Repo: repo,
// 	}
// }

// func (usecase *TeamUsecase) Create(data teams.Core) (err error) {
// 	err = usecase.Repo.Insert(data)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
