package service

import (
	"database-example/model"
	"database-example/repo"
)

type TourPointService struct {
	TourPointRepo *repo.TourPointRepository
}

func (service *TourPointService) Create(tourPoint *model.TourPoint) error {
	err := service.TourPointRepo.Create(tourPoint)
	if err != nil {
		return err
	}
	return nil
}

func (service *TourPointService) GetAll() ([]model.TourPoint, error) {
	tourPoints, err := service.TourPointRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return tourPoints, nil
}
