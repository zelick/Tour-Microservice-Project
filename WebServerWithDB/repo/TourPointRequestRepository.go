package repo

import (
	"database-example/model"

	"gorm.io/gorm"
)

type TourPointRequestRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *TourPointRequestRepository) FindById(id string) (model.TourPointRequest, error) {
	tourPointRequest := model.TourPointRequest{}
	dbResult := repo.DatabaseConnection.First(&tourPointRequest, "id = ?", id)
	if dbResult != nil {
		return tourPointRequest, dbResult.Error
	}
	return tourPointRequest, nil
}

func (repo *TourPointRequestRepository) CreateTourPointRequest(tourPointRequest *model.TourPointRequest) error {
	dbResult := repo.DatabaseConnection.Create(tourPointRequest)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *TourPointRequestRepository) UpdateTourPointRequest(tourPointRequest *model.TourPointRequest) error {
	dbResult := repo.DatabaseConnection.Save(tourPointRequest)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	return nil
}
