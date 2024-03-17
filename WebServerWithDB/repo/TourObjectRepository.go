package repo

import (
	"database-example/model"

	"gorm.io/gorm"
)

type TourObjectRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *TourObjectRepository) FindById(id string) (model.TourObject, error) {
	tourObject := model.TourObject{}
	dbResult := repo.DatabaseConnection.First(&tourObject, "id = ?", id)
	if dbResult != nil {
		return tourObject, dbResult.Error
	}
	return tourObject, nil
}

func (repo *TourObjectRepository) CreateTourObject(tourObject *model.TourObject) error {
	dbResult := repo.DatabaseConnection.Create(tourObject)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}
