// Package controllers ...
package controllers

import (
	"server/config"
	"server/models"

	"github.com/gin-gonic/gin"
)

// UpdateWallet ..
func UpdateWallet(c *gin.Context) {
	type DataType struct {
		Balance    float64           `json:"balance"`
		UserID     float64           `json:"userID"`
		WalletLogs models.WalletLogs `json:"walletLogs"`
	}

	var data DataType
	c.ShouldBindJSON(&data)

	// First get the wallet
	var wallet models.Wallet
	config.DB.Where("user_id = ?", data.UserID).First(&wallet)
	// incress balance
	if data.WalletLogs.Type == "increase" {
		wallet.Balance = wallet.Balance + data.Balance
	}
	if data.WalletLogs.Type == "decrease" {
		wallet.Balance = wallet.Balance - data.Balance
	}
	// Save
	config.DB.Save(&wallet)

	// Add wallet logs
	config.DB.Create(&data.WalletLogs)

	c.JSON(200, gin.H{
		"message": "Success",
	})

}
