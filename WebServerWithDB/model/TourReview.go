package model

import (
	"time"
)

type TourReview struct {
	ID             int       `json:"id" gorm:"column:Id;primaryKey;autoIncrement"`
	TourID         int       `json:"tourId"`
	Grade          float64   `json:"grade"`
	Comment        string    `json:"comment"`
	TouristID      int       `json:"touristId"` //ovde izmena?
	AttendanceDate time.Time `json:"attendanceDate"`
	ReviewDate     time.Time `json:"reviewDate"`
	Images         []string  `json:"images" gorm:"type:varchar[]"` //upis u bazi- izmena?
}
