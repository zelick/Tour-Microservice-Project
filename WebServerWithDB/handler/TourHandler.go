package handler

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type TourHandler struct {
	TourService *service.TourService
}

func (handler *TourHandler) Get(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	log.Printf("Tour sa id-em %s", id)
	// student, err := handler.StudentService.FindStudent(id)
	// writer.Header().Set("Content-Type", "application/json")
	// if err != nil {
	// 	writer.WriteHeader(http.StatusNotFound)
	// 	return
	// }
	writer.WriteHeader(http.StatusOK)
	// json.NewEncoder(writer).Encode(student)
}

func (handler *TourHandler) Create(writer http.ResponseWriter, req *http.Request) {
	//ResponseWriter - pisanje odgovora
	//Request - dolazni zahtev
	var tour model.Tour
	err := json.NewDecoder(req.Body).Decode(&tour) //dekodiranje json zahteva
	if err != nil {
		println("Error while parsing json")
		println("Greska:", err.Error())
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.TourService.Create(&tour)
	if err != nil {
		println("Error while creating a new tour")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(tour) // dodala sam
}

func (handler *TourHandler) GetByUserId(writer http.ResponseWriter, req *http.Request) {
	// Koristimo mux.Vars da bismo dobili vrednosti putanje
	vars := mux.Vars(req)
	userIdStr, ok := vars["userId"]
	if !ok {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	// Konvertujemo userID iz stringa u integer
	userID, err := strconv.Atoi(userIdStr)
	if err != nil {
		println(err.Error())
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	tours, err := handler.TourService.FindByUserId(userID)
	if err != nil {
		println(err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(tours)
	if err != nil {
		println(err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Postavljanje Content-Type zaglavlja odgovora
	writer.Header().Set("Content-Type", "application/json")
	// Slanje odgovora sa status kodom 200 i JSON tura
	writer.WriteHeader(http.StatusOK)
	writer.Write(jsonResponse)
}

func (handler *TourHandler) AddCharacteristics(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	tourIDStr, ok := vars["id"]
	if !ok {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	tourID, err := strconv.Atoi(tourIDStr)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	var characteristic model.TourCharacteristic
	err = json.NewDecoder(req.Body).Decode(&characteristic)
	if err != nil {
		log.Printf("Error while parsing JSON: %s", err.Error())
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.TourService.SetTourCharacteristic(tourID, characteristic.Distance, characteristic.Duration, characteristic.TransportType)
	if err != nil {
		log.Printf("Error while setting tour characteristics: %s", err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
}
