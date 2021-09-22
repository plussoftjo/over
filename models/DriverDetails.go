// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// DriverDetails ..
type DriverDetails struct {
	gorm.Model
	UserID    uint    `json:"userID"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Heading   float64 `json:"heading"`
	Status    int64   `json:"status"` // 01 => offline, 02 => online
}
