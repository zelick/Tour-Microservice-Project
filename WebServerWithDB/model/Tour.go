package model

import (
	"time"

	"github.com/lib/pq"
)

type TourStatus int

const (
	Draft TourStatus = iota
	Published
	Archived
)

type DifficultyLevel int

const (
	Easy DifficultyLevel = iota
	Moderate
	Difficult
)

type Tour struct {
	ID                  int                  `json:"id" gorm:"column:Id;primaryKey;autoIncrement"`
	Name                string               `json:"name" gorm:"not null;type:string"`
	DifficultyLevel     string               `json:"difficultyLevel"`
	Description         string               `json:"description"`
	Tags                pq.StringArray       `json:"tags" gorm:"type:text[]"`
	Status              string               `json:"status"`
	Price               int                  `json:"price"`
	UserID              int                  `json:"userId"`
	PublishedDateTime   time.Time            `json:"publishedDateTime"`
	ArchivedDateTime    time.Time            `json:"archivedDateTime"`
	TourPoints          []TourPoint          `json:"tourPoints" gorm:"foreignKey:TourID"`
	TourCharacteristics []TourCharacteristic `json:"tourCharacteristics" gorm:"type:json"` //jer je value object
	TourReviews         []TourReview         `json:"tourReviews" gorm:"foreignKey:TourID"`
}
