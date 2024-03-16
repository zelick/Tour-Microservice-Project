package repo

import (
	"database-example/model"

	"gorm.io/gorm"
)

type TourRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *TourRepository) FindById(id string) (model.Tour, error) {
	tour := model.Tour{}
	dbResult := repo.DatabaseConnection.First(&tour, "id = ?", id)
	if dbResult != nil {
		return tour, dbResult.Error
	}
	return tour, nil
}

func (repo *TourRepository) CreateTour(tour *model.Tour) error {
	dbResult := repo.DatabaseConnection.Create(tour)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *TourRepository) FindByUserId(userID int) ([]model.Tour, error) {
	var tours []model.Tour
	dbResult := repo.DatabaseConnection.Where("user_id = ?", userID).Find(&tours) //autorId
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return tours, nil
}
