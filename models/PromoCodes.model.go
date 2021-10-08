// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// PromoCodes ..
type PromoCodes struct {
	gorm.Model
	Title string  `json:"title"`
	Code  string  `json:"code"`
	Type  string  `json:"type"` // discount, credit
	Value float64 `json:"value"`
}
