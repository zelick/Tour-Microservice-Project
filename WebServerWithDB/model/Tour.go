package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
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
	ID                int            `json:"id" gorm:"column:Id;primaryKey;autoIncrement"`
	Name              string         `json:"name" gorm:"not null;type:string"`
	DifficultyLevel   string         `json:"difficultyLevel"`
	Description       string         `json:"description"`
	Tags              pq.StringArray `json:"tags" gorm:"type:text[]"`
	Status            string         `json:"status"`
	Price             int            `json:"price"`
	UserID            int            `json:"userId"`
	PublishedDateTime time.Time      `json:"publishedDateTime"`
	ArchivedDateTime  time.Time      `json:"archivedDateTime"`
	TourPoints        []TourPoint    `json:"tourPoints" gorm:"foreignKey:TourID"`
	//TourCharacteristics []TourCharacteristic `json:"tourCharacteristics" gorm:"type:jsonb"` //jer je value object
	TourCharacteristics TourCharacteristicsSlice `json:"tourCharacteristics" gorm:"type:jsonb"`
	TourReviews         []TourReview             `json:"tourReviews" gorm:"foreignKey:TourID"`
}

type TourCharacteristic struct {
	//ID            int     `json:"id" gorm:"column:Id;primaryKey;autoIncrement"` //kako kad je value objet - izmeniti
	Distance      float64 `json:"distance"`
	Duration      float64 `json:"duration"`
	TransportType string  `json:"transportType"`
}

type TransportType int

const (
	Walking TransportType = iota
	Biking
	Driving
)

/*
	func (tc TourCharacteristic) Value() (driver.Value, error) {
		return json.Marshal(tc)
	}
*/
func (tcs TourCharacteristicsSlice) Value() (driver.Value, error) {
	jsonBytes, err := json.Marshal(tcs)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal TourCharacteristicsSlice to JSON: %v", err)
	}
	return jsonBytes, nil
}

type TourCharacteristicsSlice []TourCharacteristic

func (tcs *TourCharacteristicsSlice) Scan(src interface{}) error {
	switch v := src.(type) {
	case []byte:
		return json.Unmarshal(v, tcs)
	case string:
		return json.Unmarshal([]byte(v), tcs)
	}
	return fmt.Errorf("type assertion failed")
}
