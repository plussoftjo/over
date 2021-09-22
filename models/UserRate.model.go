// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// UserRate ..
type UserRate struct {
	gorm.Model
	UserID  uint    `json:"userID"`
	RaterID uint    `json:"raterID"`
	OrderID uint    `json:"orderID"`
	Note    string  `json:"note"`
	Stars   float64 `json:"stars"`
}
