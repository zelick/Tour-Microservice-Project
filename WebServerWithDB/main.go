package main

import (
	"database-example/handler"
	"database-example/model"
	"database-example/repo"
	"database-example/service"
	"log"
	"net/http"

	"gorm.io/driver/postgres"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	dsn := "user=postgres password=super dbname=explorer-v1 host=localhost port=5432 sslmode=disable search_path=tours" // podesavanje baze
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		print(err)
		return nil
	}

	database.AutoMigrate(&model.Tour{}, &model.TourPoint{}, &model.TourReview{}, &model.TourObject{}) // migracije da bismo napravili tabele
	//database.Exec("INSERT IGNORE INTO students VALUES ('aec7e123-233d-4a09-a289-75308ea5b7e6', 'Marko Markovic', 'Graficki dizajn')")
	return database
}

func startTourServer(handler *handler.TourHandler, tourObjectHandler *handler.TourObjectHandler) {
	router := mux.NewRouter().StrictSlash(true)

	//router.HandleFunc("/students/{id}", handler.Get).Methods("GET")
	//router.HandleFunc("/students", handler.Create).Methods("POST")

	// tours
	router.HandleFunc("/tours/{id}", handler.Get).Methods("GET")
	router.HandleFunc("/tours/create", handler.Create).Methods("POST")
	router.HandleFunc("/tours/getByAuthor/{userId}", handler.GetByUserId).Methods("GET")
	router.HandleFunc("/tours/publish/{tourId}", handler.Publish).Methods("PUT")
	router.HandleFunc("/tours/archive/{tourId}", handler.Archive).Methods("PUT")
	router.HandleFunc("/tours/delete/{tourId}", handler.Delete).Methods("DELETE")

	// tour objects
	router.HandleFunc("/tourObjects/{id}", tourObjectHandler.Get).Methods("GET")
	router.HandleFunc("/tourObjects/create", tourObjectHandler.Create).Methods("POST")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	println("Server starting")
	//log.Fatal(http.ListenAndServe(":88", router))
	log.Fatal(http.ListenAndServe(":3000", router))
}

func main() {
	database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}
	tourRepo := &repo.TourRepository{DatabaseConnection: database}
	tourService := &service.TourService{TourRepo: tourRepo}
	tourHandler := &handler.TourHandler{TourService: tourService}

	tourObjectRepo := &repo.TourObjectRepository{DatabaseConnection: database}
	tourObjectService := &service.TourObjectService{TourObjectRepo: tourObjectRepo}
	tourObjectHandler := &handler.TourObjectHandler{TourObjectService: tourObjectService}

	startTourServer(tourHandler, tourObjectHandler)
}
