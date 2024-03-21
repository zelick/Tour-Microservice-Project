package handler

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type TourPointHandler struct {
	TourPointService *service.TourPointService
}

func (handler *TourPointHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var tourPoint model.TourPoint
	err := json.NewDecoder(req.Body).Decode(&tourPoint)
	if err != nil {
		println("Error while parsing json")
		println("Greska:", err.Error())
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.TourPointService.Create(&tourPoint)
	if err != nil {
		println("Error while creating a new tour point")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(tourPoint) // dodala sam
}

func (handler *TourPointHandler) GetAll(writer http.ResponseWriter, req *http.Request) {
	tourPoints, err := handler.TourPointService.GetAll()

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(writer, "Failed to get tour points: %v", err)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(tourPoints); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(writer, "Failed to encode tour points to JSON: %v", err)
		return
	}
}

func (handler *TourPointHandler) GetById(writer http.ResponseWriter, req *http.Request) {
	// Izvlačenje ID-a tačke ture iz putanje
	vars := mux.Vars(req)
	id, ok := vars["id"]
	if !ok {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(writer, "ID parameter is missing")
		return
	}

	// Konvertovanje ID-a u tip int
	tourPointId, err := strconv.Atoi(id)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(writer, "Invalid tour point ID: %v", err)
		return
	}

	// Dobijanje tačke ture po ID-u
	tourPoint, err := handler.TourPointService.FindById(tourPointId)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(writer, "Tour point not found: %v", err)
		return
	}

	// Slanje odgovora sa podacima o tački ture
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(tourPoint); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(writer, "Failed to encode tour point to JSON: %v", err)
		return
	}
}
