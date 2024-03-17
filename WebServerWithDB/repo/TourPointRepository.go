package repo

import (
	"database-example/model"

	"gorm.io/gorm"
)

type TourPointRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *TourPointRepository) Create(tourPoint *model.TourPoint) error {
	dbResult := repo.DatabaseConnection.Create(tourPoint)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}
