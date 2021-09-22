// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// WalletLogs ..
type WalletLogs struct {
	gorm.Model
	UserID     uint    `json:"userID"`
	Balance    float64 `json:"balance" gorm:"default:0"`
	OldBalance float64 `json:"oldBalance"`
	Type       string  `json:"type"` // Incress => incress from the payment, Decress => decres from the wallet
	OrderID    uint    `json:"orderID" gorm:"default:0"`
	PaymentID  uint    `json:"paymentID" gorm:"default:0"`
}
