package repo

import (
	"database-example/model"
	"errors"
	"fmt"

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

func (repo *TourRepository) UpdateTour(tour *model.Tour) error {
	dbResult := repo.DatabaseConnection.Save(tour)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	return nil
}

func (repo *TourRepository) GetTourById(id int) (model.Tour, error) {
	var tour model.Tour
	dbResult := repo.DatabaseConnection.First(&tour, id)
	if dbResult.Error != nil {
		return model.Tour{}, dbResult.Error
	}
	return tour, nil
}

/*func (repo *TourRepository) DeleteTour(id int) error {
	dbResult := repo.DatabaseConnection.Delete(&model.Tour{}, id)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	return nil
}*/

func (repo *TourRepository) DeleteTour(tourID int) error {
	// Find the tour to delete
	var tour model.Tour
	result := repo.DatabaseConnection.Preload("TourPoints").Preload("TourReviews").First(&tour, tourID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return fmt.Errorf("tour with ID %d not found", tourID)
		}
		return result.Error
	}

	err := repo.DatabaseConnection.Where("tour_id = ?", tourID).Delete(&model.TourPoint{}).Error
	if err != nil {
		println(err.Error())
		return err
	}
	err = repo.DatabaseConnection.Where("tour_id = ?", tourID).Delete(&model.TourReview{}).Error
	if err != nil {
		println(err.Error())
		return err
	}

	// Delete the tour
	err = repo.DatabaseConnection.Delete(&tour).Error
	if err != nil {
		println(err.Error())
		return err
	}

	return nil
}
