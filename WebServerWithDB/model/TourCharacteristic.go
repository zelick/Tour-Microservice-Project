package model

//"github.com/google/uuid"

type TourCharacteristic struct {
	ID            int           `json:"id" gorm:"column:Id;primaryKey;autoIncrement"`
	Distance      float64       `json:"distance"`
	Duration      float64       `json:"duration"`
	TransportType TransportType `json:"transportType"`
}

type TransportType int

const (
	Walking TransportType = iota
	Biking
	Driving
)
