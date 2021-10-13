// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// Services ..
type Services struct {
	gorm.Model
	Title           string  `json:"title"`
	Image           string  `json:"image"`
	Description     string  `json:"description"`
	Fare            float64 `json:"fare"`
	DurationFare    float64 `json:"durationFare"`
	DestinationFare float64 `json:"destinationFare"`
	WaitingTimeFare float64 `json:"waitingTimeFare"`
	Commission      float64 `json:"commission" gorm:"default:0"`
}
