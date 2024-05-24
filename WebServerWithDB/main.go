package main

import (
	"context"
	"database-example/handler"
	"database-example/model"
	"database-example/proto/tour"
	"database-example/repo"
	"database-example/service"
	"log"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	dsn := "user=postgres password=super dbname=explorer host=database port=5432 sslmode=disable search_path=tours"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		print(err)
		return nil
	}

	database.AutoMigrate(&model.Tour{}, &model.TourPoint{}, &model.TourReview{}, &model.TourObject{}, &model.PublicTourPoint{})
	return database
}

func startTourServer(handler *handler.TourHandler, tourObjectHandler *handler.TourObjectHandler, tourPointHandler *handler.TourPointHandler,
	tourPointRequestHandler *handler.TourPointRequestHandler, publicTourPointHandler *handler.PublicTourPointHandler) {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/tours/{id}", handler.Get).Methods("GET")
	router.HandleFunc("/tours/create", handler.Create).Methods("POST")
	router.HandleFunc("/tours/getByAuthor/{userId}", handler.GetByUserId).Methods("GET")
	router.HandleFunc("/tours/setCaracteristics/{id}", handler.AddCharacteristics).Methods("PUT")
	router.HandleFunc("/tours/publish/{tourId}", handler.Publish).Methods("PUT")
	router.HandleFunc("/tours/archive/{tourId}", handler.Archive).Methods("PUT")
	router.HandleFunc("/tours/delete/{tourId}", handler.Delete).Methods("DELETE")

	router.HandleFunc("/tourPoint/create", tourPointHandler.Create).Methods("POST")
	router.HandleFunc("/tourPoint/getAll", tourPointHandler.GetAll).Methods("GET")
	router.HandleFunc("/tourPoint/getById/{id}", tourPointHandler.GetById).Methods("GET")

	router.HandleFunc("/tourObjects/{id}", tourObjectHandler.Get).Methods("GET")
	router.HandleFunc("/tourObjects/create", tourObjectHandler.Create).Methods("POST")

	router.HandleFunc("/tourPointRequest/create", tourPointRequestHandler.Create).Methods("POST")
	router.HandleFunc("/tourPointRequest/accept/{tourPointRequestId}", tourPointRequestHandler.AcceptRequest).Methods("PUT")
	router.HandleFunc("/tourPointRequest/decline/{tourPointRequestId}", tourPointRequestHandler.DeclineRequest).Methods("PUT")

	router.HandleFunc("/publicTourPoint/setPublicTourPoint/{tourPointId}", publicTourPointHandler.CreateFromTourPointId).Methods("GET")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	log.Println("Server starting on port 3000")

	log.Fatal(http.ListenAndServe(":3000", router))
}

func main() {
	database := initDB()
	if database == nil {
		log.Fatal("Failed to connect to the database")
		return
	}

	tourRepo := &repo.TourRepository{DatabaseConnection: database}
	tourService := &service.TourService{TourRepo: tourRepo}

	lis, err := net.Listen("tcp", "tours:3000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	tour.RegisterTourServer(grpcServer, &Server{TourService: tourService})

	reflection.Register(grpcServer)
	log.Println("gRPC server starting on port 3000")
	grpcServer.Serve(lis)
}

type Server struct {
	tour.UnimplementedTourServer
	TourService *service.TourService
}

func (s *Server) Create(ctx context.Context, request *tour.TourDto) (*tour.TourDto, error) {
	// Map TourCharacteristicDto to model.TourCharacteristic
	println("Usao je ovde CREATE TOURS")
	var characteristics model.TourCharacteristicsSlice
	for _, c := range request.TourCharacteristics {
		characteristics = append(characteristics, model.TourCharacteristic{
			Distance:      c.Distance,
			Duration:      c.Duration,
			TransportType: c.TransportType,
		})
	}

	t := model.Tour{
		Name:                request.Name,
		Description:         request.Description,
		UserID:              int(request.UserId),
		DifficultyLevel:     request.DifficultyLevel,
		Tags:                request.Tags,
		Status:              request.Status,
		Price:               int(request.Price),
		PublishedDateTime:   request.PublishedDateTime.AsTime(),
		ArchivedDateTime:    request.ArchivedDateTime.AsTime(),
		TourCharacteristics: characteristics,
	}

	err := s.TourService.Create(&t)
	if err != nil {
		return nil, err
	}

	response := &tour.TourDto{
		Id:              int64(t.ID),
		Name:            t.Name,
		Description:     t.Description,
		DifficultyLevel: t.DifficultyLevel,
		Tags:            t.Tags,
		Status:          t.Status,
		Price:           int32(t.Price),
		UserId:          int64(t.UserID),
	}

	return response, nil
}

func (s *Server) GetByUserId(ctx context.Context, request *tour.PageRequest) (*tour.TourListResponse, error) {
	// Poziv metode FindByUserId iz TourService-a
	print("Usao je tours!")
	tours, err := s.TourService.FindByUserId(int(request.UserId))
	if err != nil {
		return nil, err
	}

	// Mapiranje tours u listu TourDto objekata
	var tourDtos []*tour.TourDto
	for _, t := range tours {
		tourDto := &tour.TourDto{
			Id:                int64(t.ID),
			Name:              t.Name,
			PublishedDateTime: timestamppb.New(t.PublishedDateTime),
			ArchivedDateTime:  timestamppb.New(t.ArchivedDateTime),
			Description:       t.Description,
			DifficultyLevel:   t.DifficultyLevel,
			Tags:              t.Tags,
			Price:             int32(t.Price),
			Status:            t.Status,
			UserId:            int64(t.UserID),
		}
		tourDtos = append(tourDtos, tourDto)
	}

	// Kreiranje TourListResponse objekta
	response := &tour.TourListResponse{
		Results:    tourDtos,
		TotalCount: int32(len(tours)),
	}

	return response, nil
}
