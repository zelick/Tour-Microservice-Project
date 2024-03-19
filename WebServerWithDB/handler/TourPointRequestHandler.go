package handler

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type TourPointRequestHandler struct {
	TourPointRequestService *service.TourPointRequestService
}

func (handler *TourPointRequestHandler) Get(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	log.Printf("Tour point request sa id-em %s", id)
	// student, err := handler.StudentService.FindStudent(id)
	// writer.Header().Set("Content-Type", "application/json")
	// if err != nil {
	// 	writer.WriteHeader(http.StatusNotFound)
	// 	return
	// }
	writer.WriteHeader(http.StatusOK)
	// json.NewEncoder(writer).Encode(student)
}

func (handler *TourPointRequestHandler) Create(writer http.ResponseWriter, req *http.Request) {
	//ResponseWriter - pisanje odgovora
	//Request - dolazni zahtev
	var request model.TourPointRequest
	err := json.NewDecoder(req.Body).Decode(&request) //dekodiranje json zahteva
	if err != nil {
		println("Error while parsing json")
		println("Greska:", err.Error())
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.TourPointRequestService.Create(&request)
	if err != nil {
		println("Error while creating a new request")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(request) // dodala sam
}

func (handler *TourPointRequestHandler) AcceptRequest(writer http.ResponseWriter, req *http.Request) {
	tourPointRequestId := mux.Vars(req)["tourPointRequestId"]
	_, err := handler.TourPointRequestService.AcceptRequest(tourPointRequestId)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
}

func (handler *TourPointRequestHandler) DeclineRequest(writer http.ResponseWriter, req *http.Request) {
	tourPointRequestId := mux.Vars(req)["tourPointRequestId"]
	_, err := handler.TourPointRequestService.DeclineRequest(tourPointRequestId)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
}
