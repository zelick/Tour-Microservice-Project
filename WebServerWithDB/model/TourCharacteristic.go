package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TourCharacteristic struct {
	ID            uuid.UUID     `json:"id"`
	Distance      float64       `json:"distance"`
	Duration      float64       `json:"duration"`
	TransportType TransportType `json:"transportType"`
}

func (tc *TourCharacteristic) BeforeCreate(tx *gorm.DB) (err error) {
	tc.ID = uuid.New()
	return nil
}

type TransportType int

const (
	Walking TransportType = iota
	Biking
	Driving
)
