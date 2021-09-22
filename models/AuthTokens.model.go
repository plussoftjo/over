// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// AuthTokens ..
type AuthTokens struct {
	gorm.Model
	Token     string `json:"token"`
	UserID    uint   `json:"user_id"`
	ClientID  uint   `json:"client_id"`
	ExpiresAt string `json:"expires_at"`
}
