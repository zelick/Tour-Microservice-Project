package repo

import (
	"database-example/model"

	"gorm.io/gorm"
)

type PublicTourPointRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *PublicTourPointRepository) Create(publicTourPoint *model.PublicTourPoint) error {
	dbResult := repo.DatabaseConnection.Create(publicTourPoint)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}
