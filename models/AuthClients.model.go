// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// AuthClients ..
type AuthClients struct {
	gorm.Model
	Name   string `json:"name"`
	Secret string `json:"secret"`
	Active int    `json:"Active"`
}
