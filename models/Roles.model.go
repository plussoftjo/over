// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// Roles ..
type Roles struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Pages       string `json:"pages"`
	Scopes      string `json:"scopes"`
}
