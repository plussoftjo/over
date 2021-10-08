// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// UserPromoCodes ..
type UserPromoCodes struct {
	gorm.Model
	UserID      uint    `json:"userID"`
	PromoCodeID uint    `json:"promoCode"`
	Title       string  `json:"title"`
	Code        string  `json:"code"`
	Value       float64 `json:"value"`
	Type        string  `json:"type"`
	Used        int     `json:"used" gorm:"default:0"`
}
