// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// DriverDetails ..
type DriverDetails struct {
	gorm.Model
	UserID    uint    `json:"userID"`
	Latitude  float64 `json:"latitude" gorm:"default:0"`
	Longitude float64 `json:"longitude" gorm:"default:0"`
	Heading   float64 `json:"heading" gorm:"default:0"`
	Status    int64   `json:"status" gorm:"default:0"` // 01 => offline, 02 => online, 03 => busy
}
