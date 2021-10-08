// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// NotificationTokens ..
type NotificationTokens struct {
	gorm.Model
	UserID uint   `json:"userID"`
	Token  string `json:"token"`
	Active int    `json:"active" gorm:"default:1"`
}
