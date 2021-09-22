// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// RiderLocations ..
type RiderLocations struct {
	gorm.Model
	Title     string  `json:"title"`
	Geo       string  `json:"geo"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	UserID    uint64  `json:"userID"`
}
