package service

import (
	"database-example/model"
	"database-example/repo"
	"fmt"
)

type TourObjectService struct {
	TourObjectRepo *repo.TourObjectRepository
}

func (service *TourObjectService) FindTourObject(id string) (*model.TourObject, error) {
	tour, err := service.TourObjectRepo.FindById(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("menu item with id %s not found", id))
	}
	return &tour, nil
}

func (service *TourObjectService) Create(tourObject *model.TourObject) error {
	err := service.TourObjectRepo.CreateTourObject(tourObject)
	if err != nil {
		return err
	}
	return nil
}
