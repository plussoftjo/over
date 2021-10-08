// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// DriversCar ..
type DriversCar struct {
	gorm.Model
	UserID   uint   `json:"userID"`
	CarBrand string `json:"carBrand"`
	CarModel string `json:"carModel"`
	CarYear  string `json:"carYear"`
	CarColor string `json:"carColor"`
	CarPlate string `json:"carPlate"`
	CarImage string `json:"carImage"`
}
