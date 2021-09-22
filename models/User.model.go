// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// DriverWithDetails ..
func DriverWithDetails(db *gorm.DB) *gorm.DB {
	return db.Preload("DriverDetail").Preload("DriverCar").Preload("Wallet")
}

// User ..
type User struct {
	gorm.Model
	Name         string        `json:"name"`
	Phone        string        `json:"phone" gorm:"unique"`
	Password     string        `json:"password"`
	RolesID      uint          `json:"roles_id"`
	UserType     uint          `json:"user_type"` // 01 -> User , 02 -> Supplier, 03 -> Controller
	DriverDetail DriverDetails `json:"driverDetail" gorm:"foreignKey:UserID;references:ID"`
	DriverCar    DriversCar    `json:"driverCar" gorm:"foreignKey:UserID;references:ID"`
	Roles        Roles         `json:"roles" gorm:"foreignKey:RolesID;references:ID"`
	Wallet       Wallet        `json:"wallet" gorm:"foreignKey:UserID;references:ID"`
}

// Login ...
type Login struct {
	Phone    string `json:"phone" gorm:"unique" binding:"required"`
	Password string `json:"password" binding:"required"`
}
