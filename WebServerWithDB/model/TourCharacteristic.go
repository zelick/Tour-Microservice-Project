package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type TourCharacteristic struct {
	ID            int     `json:"id" gorm:"column:Id;primaryKey;autoIncrement"` //kako kad je value objet - izmeniti
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

func (tc TourCharacteristic) Value() (driver.Value, error) {
	return json.Marshal(tc)
}

func (tc *TourCharacteristic) Scan(value interface{}) error {
	if value == nil {
		*tc = TourCharacteristic{}
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("Scan source is not []byte")
	}
	return json.Unmarshal(bytes, tc)
}
