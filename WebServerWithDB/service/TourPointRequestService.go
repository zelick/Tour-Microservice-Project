package service

import (
	"database-example/model"
	"database-example/repo"
	"fmt"
)

type TourPointRequestService struct {
	TourPointRequestRepo *repo.TourPointRequestRepository
}

func (service *TourPointRequestService) FindTourPointRequest(id string) (*model.TourPointRequest, error) { //ovo se ne koristi
	tour, err := service.TourPointRequestRepo.FindById(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("menu item with id %s not found", id))
	}
	return &tour, nil
}

func (service *TourPointRequestService) Create(TourPointRequest *model.TourPointRequest) error {
	err := service.TourPointRequestRepo.CreateTourPointRequest(TourPointRequest)
	if err != nil {
		return err
	}
	return nil
}

func (service *TourPointRequestService) AcceptRequest(tourPointRequestId string) (interface{}, error) {
	request, err := service.TourPointRequestRepo.FindById(tourPointRequestId)
	if err != nil {
		return nil, err
	}
	request.Status = "Accepted"
	err = service.TourPointRequestRepo.UpdateTourPointRequest(&request)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (service *TourPointRequestService) DeclineRequest(tourPointRequestId string) (interface{}, error) {
	request, err := service.TourPointRequestRepo.FindById(tourPointRequestId)
	if err != nil {
		return nil, err
	}
	request.Status = "Declined"
	err = service.TourPointRequestRepo.UpdateTourPointRequest(&request)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
