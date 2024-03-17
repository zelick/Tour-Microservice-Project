package model

type TourObject struct {
	ID          int     `json:"id" gorm:"column:Id;primaryKey;autoIncrement"`
	Name        string  `json:"name" gorm:"not null;type:string"`
	Description string  `json:"description"`
	ImageUrl    string  `json:"imageUrl"`
	Category    string  `json:"category"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
}
