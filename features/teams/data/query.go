package data

import (
	"errors"
	"immersiveProject/features/class/data"
	"immersiveProject/features/teams"

	// _ "immersiveProject/features/teams/data"

	"gorm.io/gorm"
)

type teamsRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) *teamsRepo {
	return &teamsRepo{
		db: db,
	}
}

func (repo *teamsRepo) Insert(data teams.Core) (affectedRow int, err error) {
	var dataInsert = entityToModel(data)
	tx := repo.db.Create(&dataInsert)
	if tx.Error != nil {
		return -1, err
	}

	return int(tx.RowsAffected), nil
}

func (repo *teamsRepo) Delete(team teams.Core) (affectedRow int, err error) {
	teamModel := Team{}
	teamModel.ID = team.ID
	tx := repo.db.Model(Team{})

	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("failed to delete class")
	}
	return int(tx.RowsAffected), nil
}

func (repo *teamsRepo) Update(teams teams.Core) (affectedRow int, err error) {
	teamsModel := entityToModel(teams)
	tx := repo.db.Model(&data.Class{}).Where("id = ?", teams.ID).Updates(&teamsModel)
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("failed to update data")
	}
	return int(tx.RowsAffected), nil
}

func (repo *teamsRepo) FindAll() (result []teams.Core, err error) {
	teamsModels := []Team{}
	tx := repo.db.Model(&Team{})
	if tx.Error != nil {
		return result, tx.Error
	}

	for _, teams := range teamsModels {
		result = append(result, ModelToEntity(teams))
	}

	return result, nil
}
