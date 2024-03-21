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

	database.AutoMigrate(&model.Tour{}, &model.TourPoint{}, &model.TourReview{}, &model.TourObject{}, &model.TourPointRequest{}, &model.PublicTourPoint{}) // migracije da bismo napravili tabele
	//database.Exec("INSERT IGNORE INTO students VALUES ('aec7e123-233d-4a09-a289-75308ea5b7e6', 'Marko Markovic', 'Graficki dizajn')")
	return database
}

func startTourServer(handler *handler.TourHandler, tourObjectHandler *handler.TourObjectHandler, tourPointHandler *handler.TourPointHandler,
	tourPointRequestHandler *handler.TourPointRequestHandler, publicTourPointHandler *handler.PublicTourPointHandler) {

	router := mux.NewRouter().StrictSlash(true)

	//router.HandleFunc("/students/{id}", handler.Get).Methods("GET")
	//router.HandleFunc("/students", handler.Create).Methods("POST")

	// tours
	router.HandleFunc("/tours/{id}", handler.Get).Methods("GET")
	router.HandleFunc("/tours/create", handler.Create).Methods("POST")
	router.HandleFunc("/tours/getByAuthor/{userId}", handler.GetByUserId).Methods("GET")
	router.HandleFunc("/tours/setCaracteristics/{id}", handler.AddCharacteristics).Methods("PUT")
	router.HandleFunc("/tours/publish/{tourId}", handler.Publish).Methods("PUT")
	router.HandleFunc("/tours/archive/{tourId}", handler.Archive).Methods("PUT")
	router.HandleFunc("/tours/delete/{tourId}", handler.Delete).Methods("DELETE")

	//tour point
	router.HandleFunc("/tourPoint/create", tourPointHandler.Create).Methods("POST")
	router.HandleFunc("/tourPoint/getAll", tourPointHandler.GetAll).Methods("GET")
	router.HandleFunc("/tourPoint/getById/{id}", tourPointHandler.GetById).Methods("GET")

	// tour objects
	router.HandleFunc("/tourObjects/{id}", tourObjectHandler.Get).Methods("GET")
	router.HandleFunc("/tourObjects/create", tourObjectHandler.Create).Methods("POST")

	// tour point requests
	router.HandleFunc("/tourPointRequest/create", tourPointRequestHandler.Create).Methods("POST")
	router.HandleFunc("/tourPointRequest/accept/{tourPointRequestId}", tourPointRequestHandler.AcceptRequest).Methods("PUT")
	router.HandleFunc("/tourPointRequest/decline/{tourPointRequestId}", tourPointRequestHandler.DeclineRequest).Methods("PUT")

	//public tour point
	router.HandleFunc("/publicTourPoint/setPublicTourPoint/{tourPointId}", publicTourPointHandler.CreateFromTourPointId).Methods("GET") //izmena po potrebi, poziv stakholders modula

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

	//tour
	tourRepo := &repo.TourRepository{DatabaseConnection: database}
	tourService := &service.TourService{TourRepo: tourRepo}
	tourHandler := &handler.TourHandler{TourService: tourService}

	//tourPoint
	tourPointRepo := &repo.TourPointRepository{DatabaseConnection: database}
	tourPointService := &service.TourPointService{TourPointRepo: tourPointRepo}
	tourPointHandler := &handler.TourPointHandler{TourPointService: tourPointService}

	//tourObject
	tourObjectRepo := &repo.TourObjectRepository{DatabaseConnection: database}
	tourObjectService := &service.TourObjectService{TourObjectRepo: tourObjectRepo}
	tourObjectHandler := &handler.TourObjectHandler{TourObjectService: tourObjectService}

	//tourPointRequest - obrisati?
	tourPointRequestRepo := &repo.TourPointRequestRepository{DatabaseConnection: database}
	tourPointRequestService := &service.TourPointRequestService{TourPointRequestRepo: tourPointRequestRepo}
	tourPointRequestHandler := &handler.TourPointRequestHandler{TourPointRequestService: tourPointRequestService}

	//publicTourPoint
	publicTourPointRepo := &repo.PublicTourPointRepository{DatabaseConnection: database}
	publicTourPointService := &service.PublicTourPointService{PublicTourPointRepo: publicTourPointRepo}
	publicTourPointHandler := &handler.PublicTourPointHandler{
		PublicTourPointService: publicTourPointService,
		TourPointService:       tourPointService,
	}

	startTourServer(tourHandler, tourObjectHandler, tourPointHandler, tourPointRequestHandler, publicTourPointHandler)
}
