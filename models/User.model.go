// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// DriverWithDetails ..
func DriverWithDetails(db *gorm.DB) *gorm.DB {
	return db.Preload("DriverDetail").Preload("DriverCar").Preload("Wallet")
}

// DriverWithDashboardDetails ..
func DriverWithDashboardDetails(db *gorm.DB) *gorm.DB {
	return db.Preload("Wallet").Preload("DriverCar").Preload("DriverDetail").Preload("DriverValues").Preload("WalletLogs")
}

// User ..
type User struct {
	gorm.Model
	Name           string           `json:"name"`
	Phone          string           `json:"phone" gorm:"unique"`
	Password       string           `json:"password"`
	RolesID        uint             `json:"roles_id"`
	UserType       uint             `json:"user_type"` // 01 -> User , 02 -> Supplier, 03 -> Controller
	PhoneCode      string           `json:"phoneCode"`
	Avatar         string           `json:"avatar"`
	Block          int              `json:"block" gorm:"default:0"`          // 0 => Active, 01 => Block
	RegisterStatus int              `json:"registerStatus" gorm:"default:0"` // 00 => RegisterButNoSetTheFirstDetails, 01 => OnStepOne, 02 => OnStepTow, 100 => complete, 101 => approved, 500 => canceld
	DriverDetail   DriverDetails    `json:"driverDetails" gorm:"foreignKey:UserID;references:ID"`
	DriverCar      DriversCar       `json:"driverCar" gorm:"foreignKey:UserID;references:ID"`
	DriverValues   []DriverValues   `json:"driverValues" gorm:"foreignKey:UserID;references:ID"`
	Roles          Roles            `json:"roles" gorm:"foreignKey:RolesID;references:ID"`
	Wallet         Wallet           `json:"wallet" gorm:"foreignKey:UserID;references:ID"`
	WalletLogs     []WalletLogs     `json:"walletLogs" gorm:"foreignKey:UserID;references:ID"`
	UserPromoCodes []UserPromoCodes `json:"userPromoCodes" gorm:"foreignKey:UserID;references:ID"`
}

// Login ...
type Login struct {
	Phone    string `json:"phone" gorm:"unique" binding:"required"`
	Password string `json:"password" binding:"required"`
}
