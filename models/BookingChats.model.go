// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// BookingChats ..
type BookingChats struct {
	gorm.Model
	BookingID uint   `json:"bookingID"`
	From      string `json:"from"`
	Text      string `json:"text"`
}
