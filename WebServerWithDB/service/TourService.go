package service

import (
	"database-example/model"
	"database-example/repo"
	"fmt"
	"time"
)

type TourService struct {
	TourRepo *repo.TourRepository
}

func (service *TourService) FindTour(id string) (*model.Tour, error) { //ovo se ne koristi
	tour, err := service.TourRepo.FindById(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("menu item with id %s not found", id))
	}
	return &tour, nil
}

func (service *TourService) Create(tour *model.Tour) error {
	err := service.TourRepo.CreateTour(tour)
	if err != nil {
		return err
	}
	return nil
}

func (service *TourService) FindByUserId(userID int) ([]model.Tour, error) {
	tours, err := service.TourRepo.FindByUserId(userID)
	if err != nil {
		return nil, err
	}
	return tours, nil
}

func (service *TourService) PublishTour(tourId int) (interface{}, error) {

	tour, err := service.TourRepo.GetTourById(tourId)
	if err != nil {
		return nil, err
	}
	if tour.Name == "" || tour.Description == "" {
		return nil, fmt.Errorf("tour must have all basic data set.")
	}
	/*if len(tour.TourPoints) < 2 {
		return nil, fmt.Errorf("Tour must have at least two key points.")
	}
	validTimeDefined := false
	for _, tc := range tour.TourCharacteristics {
		if tc.Duration > 0 {
			validTimeDefined = true
			break
		}
	}
	if !validTimeDefined {
		return nil, fmt.Errorf("At least one valid tour time must be defined.")
	}*/
	tour.Status = "Published"
	err = service.TourRepo.UpdateTour(&tour)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (service *TourService) PublishTourNEW(tourId int) (*model.Tour, error) {
	tour, err := service.TourRepo.GetTourById(tourId)
	if err != nil {
		return nil, err
	}

	tour.Status = "Published"
	tour.PublishedDateTime = time.Now() // Postavljanje vremena objavljivanja
	err = service.TourRepo.UpdateTour(&tour)
	if err != nil {
		return nil, err
	}

	return &tour, nil
}

func (service *TourService) ArchiveTourNEW(tourId int) (*model.Tour, error) {
	tour, err := service.TourRepo.GetTourById(tourId)
	if err != nil {
		return nil, err
	}
	tour.Status = "Archived"
	tour.ArchivedDateTime = time.Now() // Postavljanje vremena arhiviranja
	err = service.TourRepo.UpdateTour(&tour)
	if err != nil {
		return nil, err
	}

	return &tour, nil
}

func (service *TourService) SetTourCharacteristic(tourID int, distance, duration float64, transportType string) error {
	tour, err := service.TourRepo.GetTourById(tourID)
	if err != nil {
		return fmt.Errorf("tour with ID %d not found", tourID)
	}

	newCharacteristic := model.TourCharacteristic{
		Distance:      distance,
		Duration:      duration,
		TransportType: transportType,
	}

	// Provera da li je TourCharacteristicsSlice prazan i inicijalizacija ako jeste
	if tour.TourCharacteristics == nil {
		tour.TourCharacteristics = make(model.TourCharacteristicsSlice, 0)
	}

	// Dodavanje novog JSONB objekta u niz
	tour.TourCharacteristics = append(tour.TourCharacteristics, newCharacteristic)

	// Ažuriranje tura u repozitorijumu
	if err := service.TourRepo.UpdateTour(&tour); err != nil {
		return fmt.Errorf("failed to update tour: %v", err)
	}
	return nil
}

func (service *TourService) DeleteTour(tourId int) error {
	err := service.TourRepo.DeleteTour(tourId)
	if err != nil {
		return err
	}
	return nil
}

func (service *TourService) ArchiveTour(tourId int) (interface{}, error) {
	tour, err := service.TourRepo.GetTourById(tourId)
	if err != nil {
		return nil, err
	}
	tour.Status = "Archived"
	err = service.TourRepo.UpdateTour(&tour)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
