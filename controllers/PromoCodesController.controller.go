// Package controllers ...
package controllers

import (
	"fmt"
	"server/config"
	"server/models"
	"time"

	"net/http"

	"github.com/gin-gonic/gin"
)

// ------------- PromoCode -------------//

// StorePromoCode ..
func StorePromoCode(c *gin.Context) {
	// Register var and bind json
	var data models.PromoCodes
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// StoreInDB
	if err := config.DB.Create(&data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, data)
}

// IndexPromoCodes ..
func IndexPromoCodes(c *gin.Context) {
	// Register PromoCodes in var
	var data []models.PromoCodes
	config.DB.Find(&data)
	c.JSON(200, data)
}

// DestroyPromoCode ..
func DestroyPromoCode(c *gin.Context) {
	ID := c.Param("id")
	config.DB.Delete(&models.PromoCodes{}, ID)
	// Return Promos list
	var data []models.PromoCodes
	config.DB.Find(&data)
	c.JSON(200, data)
}

// UpdatePromoCode ..
func UpdatePromoCode(c *gin.Context) {
	var promoCode models.PromoCodes
	if err := c.ShouldBindJSON(&promoCode); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Model(&promoCode).Update(&promoCode).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var promoCodes []models.PromoCodes
	config.DB.Find(&promoCodes)
	c.JSON(200, gin.H{
		"promoCode":  promoCode,
		"promoCodes": promoCodes,
	})
}

// ------------- End Promo Code ------------ //

// CheckPromoCode ..
func CheckPromoCode(c *gin.Context) {
	type CheckPromoCodeProps struct {
		UserID uint   `json:"userID"`
		Code   string `json:"code"`
	}
	var RequestProps CheckPromoCodeProps
	c.ShouldBindJSON(&RequestProps)

	go func() {
		time.Sleep(time.Second * 5)
		fmt.Println("After 20 second call ID: ")
	}()

	/**
	* Codes
	 */
	var promoCode models.PromoCodes
	pErr := config.DB.Where("code = ?", RequestProps.Code).First(&promoCode).Error
	if pErr != nil {
		c.JSON(500, gin.H{
			"err":     pErr.Error(),
			"code":    101,
			"message": "No code valid",
		})
		return
	}

	var userPromoCode int64
	config.DB.Model(&models.UserPromoCodes{}).Where("user_id = ?", RequestProps.UserID).Where("promo_code_id = ?", promoCode.ID).Count(&userPromoCode)
	if userPromoCode != 0 {
		c.JSON(500, gin.H{
			"code":    102,
			"message": "Code used",
		})
		return
	}
	config.DB.Create(&models.UserPromoCodes{
		UserID:      RequestProps.UserID,
		PromoCodeID: promoCode.ID,
		Title:       promoCode.Title,
		Value:       promoCode.Value,
		Code:        promoCode.Code,
		Type:        promoCode.Type,
		Used:        0,
	})

	if promoCode.Type == "credit" {
		var walletBalance models.Wallet
		err := config.DB.Model(&models.Wallet{}).Where("user_id = ?", RequestProps.UserID).First(&walletBalance).Error
		if err != nil {
			c.JSON(500, gin.H{
				"err":  err.Error(),
				"code": 101,
			})
			return
		}

		// Set New Wallet balance
		newWalletAmount := walletBalance.Balance + promoCode.Value
		updateErr := config.DB.Model(&models.Wallet{}).Where("user_id = ?", RequestProps.UserID).Updates(&models.Wallet{
			Balance: newWalletAmount,
		}).Error
		if updateErr != nil {
			c.JSON(500, gin.H{
				"err":  updateErr.Error(),
				"code": 101,
			})
			return
		}
		config.DB.Create(&models.WalletLogs{
			UserID:     RequestProps.UserID,
			Balance:    promoCode.Value,
			OldBalance: walletBalance.Balance,
			Type:       "increase",
			PaymentID:  promoCode.ID,
		})
	}

	c.JSON(200, gin.H{
		"code":    200,
		"message": "Added success",
	})

}
