package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TourReview struct {
	ID             uuid.UUID `json:"id"`
	TourID         uuid.UUID `json:"tourId"`
	Grade          float64   `json:"grade"`
	Comment        string    `json:"comment"`
	TouristID      int       `json:"touristId"` //ovde izmena?
	AttendanceDate time.Time `json:"attendanceDate"`
	ReviewDate     time.Time `json:"reviewDate"`
	Images         []string  `json:"images" gorm:"type:varchar[]"` //upis u bazi- izmena?
}

func (tr *TourReview) BeforeCreate(tx *gorm.DB) (err error) {
	tr.ID = uuid.New()
	return nil
}
