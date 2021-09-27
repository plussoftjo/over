// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// UserStatus ..
type UserStatus struct {
	gorm.Model
	UserID     uint `json:"userID"`
	Blocking   int  `json:"blocking"`
	Orders     int  `json:"orders"`
	CancelTrip int  `json:"cancelTrip"`
	DimssTrip  int  `json:"dimssTrip"`
}
