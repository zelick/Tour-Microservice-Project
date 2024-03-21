package service

import (
	"database-example/model"
	"database-example/repo"
)

type PublicTourPointService struct {
	PublicTourPointRepo *repo.PublicTourPointRepository
}

func (service *PublicTourPointService) Create(publicTourPoint *model.PublicTourPoint) error {
	err := service.PublicTourPointRepo.Create(publicTourPoint)
	if err != nil {
		return err
	}
	return nil
}
