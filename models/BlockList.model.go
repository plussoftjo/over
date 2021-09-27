// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// BlockList ..
type BlockList struct {
	gorm.Model
	RiderID  uint `json:"riderID"`
	DriverID uint `json:"driverID"`
}
