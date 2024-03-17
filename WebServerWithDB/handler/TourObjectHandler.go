package handler

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type TourObjectHandler struct {
	TourObjectService *service.TourObjectService
}

func (handler *TourObjectHandler) Get(writer http.ResponseWriter, req *http.Request) {
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

func (handler *TourObjectHandler) Create(writer http.ResponseWriter, req *http.Request) {
	//ResponseWriter - pisanje odgovora
	//Request - dolazni zahtev
	var tourObject model.TourObject
	err := json.NewDecoder(req.Body).Decode(&tourObject) //dekodiranje json zahteva
	if err != nil {
		println("Error while parsing json")
		println("Greska:", err.Error())
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.TourObjectService.Create(&tourObject)
	if err != nil {
		println("Error while creating a new tour")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(tourObject) // dodala sam
}
