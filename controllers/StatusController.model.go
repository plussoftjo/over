// Package controllers ...
package controllers

import (
	"server/config"
	"server/models"
	"server/vendors"

	"github.com/gin-gonic/gin"
)

// StoreUserStatus ...
func StoreUserStatus(c *gin.Context) {
	type StoreStatusType struct {
		UserID uint   `json:"userID"`
		Type   string `json:"type"`
	}
	var data StoreStatusType
	c.ShouldBindJSON(&data)

	// There are to status when its have today or not have today

	// Get today values first
	startOfDay, endOfDay := vendors.BetwenToday()

	// Register Data and check if there status today or not
	var todayRecordCount int64
	err := config.DB.Model(&models.UserStatus{}).Where("user_id = ?", data.UserID).Where("created_at BETWEEN ? AND ?", startOfDay, endOfDay).Count(&todayRecordCount).Error
	if err != nil {
		c.JSON(200, gin.H{
			"err":  err.Error(),
			"code": 101,
		})
		return
	}

	// Set the type
	blocking := 0
	orders := 0
	cancelTrip := 0
	dimssTrip := 0
	if data.Type == "blocking" {
		blocking = 1
	}
	if data.Type == "orders" {
		orders = 1
	}
	if data.Type == "cancelTrip" {
		cancelTrip = 1
	}
	if data.Type == "dimssTrip" {
		dimssTrip = 1
	}

	// Check count

	// IF not have record
	if todayRecordCount == 0 {
		// Register new record
		config.DB.Create(&models.UserStatus{
			UserID:     data.UserID,
			Blocking:   blocking,
			Orders:     orders,
			CancelTrip: cancelTrip,
			DimssTrip:  dimssTrip,
		})
	} else {
		var userStatus models.UserStatus
		config.DB.Where("user_id = ?", data.UserID).Where("created_at BETWEEN ? AND ?", startOfDay, endOfDay).First(&userStatus)

		config.DB.Model(&models.UserStatus{}).Where("user_id = ?", data.UserID).Updates(&models.UserStatus{
			Blocking:   userStatus.Blocking + blocking,
			Orders:     userStatus.Orders + orders,
			CancelTrip: userStatus.CancelTrip + cancelTrip,
			DimssTrip:  userStatus.DimssTrip + dimssTrip,
		})
	}

	c.JSON(200, gin.H{
		"message": "success",
	})

}
