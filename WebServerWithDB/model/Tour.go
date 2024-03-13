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
	TourPoints          []TourPoint          `json:"tourPoints"`
	TourCharacteristics []TourCharacteristic `gorm:"type:TourCharacteristic[]"`
	TourReviews         []TourReview         `json:"tourReviews"`
}

func (tour *Tour) BeforeCreate(scope *gorm.DB) error {
	tour.ID = uuid.New()
	return nil
}
