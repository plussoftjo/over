// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// Wallet ..
type Wallet struct {
	gorm.Model
	UserID  uint    `json:"userID"`
	Balance float64 `json:"balance" gorm:"default:0"`
}
