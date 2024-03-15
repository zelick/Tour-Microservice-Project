package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Tour struct {
	ID                  uuid.UUID            `json:"id"`
	Name                string               `json:"name" gorm:"not null;type:string"`
	DifficultyLevel     int                  `json:"difficultyLevel"`
	Description         string               `json:"description"`
	Tags                pq.StringArray       `json:"tags" gorm:"type:text[]"`
	Status              int                  `json:"status"`
	Price               int                  `json:"price"`
	UserID              int                  `json:"userId"`
	PublishedDateTime   *time.Time           `json:"publishedDateTime"`
	ArchivedDateTime    *time.Time           `json:"archivedDateTime"`
	TourPoints          []TourPoint          `json:"tourPoints" gorm:"foreignKey:TourID"`
	TourCharacteristics []TourCharacteristic `json:"tourCharacteristics" gorm:"type:json"` //jer je value object
	TourReviews         []TourReview         `json:"tourReviews" gorm:"foreignKey:TourID"`
}

func (tour *Tour) BeforeCreate(scope *gorm.DB) error {
	tour.ID = uuid.New()
	return nil
}
