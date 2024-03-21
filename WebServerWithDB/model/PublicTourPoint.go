package model

type PublicTourPoint struct {
	ID          int     `json:"id" gorm:"column:Id;primaryKey;autoIncrement"`
	IdTour      int     `json:"tourId"` //TourID
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	ImageUrl    string  `json:"imageUrl"`
	//Secret      string  `json:"secret"`
}
