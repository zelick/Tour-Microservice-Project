package handler

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type PublicTourPointHandler struct {
	PublicTourPointService *service.PublicTourPointService
	TourPointService       *service.TourPointService
}

func (handler *PublicTourPointHandler) CreateFromTourPointId(writer http.ResponseWriter, req *http.Request) {

	// Izvlačenje tourPointId iz putanje
	vars := mux.Vars(req)
	tourPointIdFromPath := vars["tourPointId"]

	// Konvertovanje tourPointId u tip koji se očekuje (int)
	tourPointId, err := strconv.Atoi(tourPointIdFromPath)
	if err != nil {
		http.Error(writer, "Invalid tour point ID", http.StatusBadRequest)
		return
	}

	// Dobijanje informacija o tački ture na osnovu njenog ID-a
	tourPoint, err := handler.TourPointService.FindById(tourPointId)
	if err != nil {
		// Greška prilikom dobijanja informacija o tački ture
		http.Error(writer, "Error while getting tour point information", http.StatusInternalServerError)
		return
	}

	// Kreiranje objekta za javnu tačku ture
	publicTourPoint := model.PublicTourPoint{
		IdTour:      tourPoint.TourID,
		Name:        tourPoint.Name,
		Description: tourPoint.Description,
		Latitude:    tourPoint.Latitude,
		Longitude:   tourPoint.Longitude,
		ImageUrl:    tourPoint.ImageURL,
	}

	// Čuvanje javne tačke ture u bazi podataka
	err = handler.PublicTourPointService.Create(&publicTourPoint)
	if err != nil {
		// Greška prilikom čuvanja javne tačke ture u bazi podataka
		http.Error(writer, "Error while creating public tour point", http.StatusInternalServerError)
		return
	}

	// Uspešno kreiranje javne tačke ture
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(publicTourPoint)
}
