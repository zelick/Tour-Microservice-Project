package handler

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"fmt"
	"net/http"
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
