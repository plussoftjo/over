// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// DriverValues ..
type DriverValues struct {
	gorm.Model
	UserID uint   `json:"userID"`
	Key    string `json:"key"`
	Value  string `json:"value"`
	Type   string `json:"type"` // text | image

}
