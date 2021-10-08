// Package controllers ...
package controllers

import (
	"server/config"
	"server/models"
	"server/vendors"

	"github.com/gin-gonic/gin"
	expo "github.com/oliveroneill/exponent-server-sdk-golang/sdk"
)

// IndexAssets ..
func IndexAssets(c *gin.Context) {
	var services []models.Services
	config.DB.Find(&services)

	c.JSON(200, gin.H{
		"services": services,
	})
}

// StoreNotificationToken ..
func StoreNotificationToken(c *gin.Context) {
	var notificationToken models.NotificationTokens
	c.ShouldBindJSON(&notificationToken)

	err := config.DB.Where("user_id = ?", notificationToken.UserID).Where("token = ?", notificationToken.Token).FirstOrCreate(&notificationToken).Error
	if err != nil {
		c.JSON(500, gin.H{
			"err":  err.Error(),
			"code": 101,
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Success",
	})
}

type JsonDataType struct {
	Type        string `json:"type"`
	BookingData string `json:"bookingData"`
}

// SendNotification .. this is global send notification
func SendNotification(c *gin.Context) {

	type notificationDataType struct {
		UserID   uint                 `json:"userID"`
		Title    string               `json:"title"`
		Body     string               `json:"body"`
		JsonData vendors.JsonDataType `json:"jsonData"`
	}

	var data notificationDataType
	c.ShouldBindJSON(&data)

	// First Get the notification token
	var token models.NotificationTokens
	config.DB.Where("user_id = ?", data.UserID).First(&token)

	var tokens []expo.ExponentPushToken
	ExpoToken, _ := expo.NewExponentPushToken(token.Token)
	tokens = append(tokens, ExpoToken)

	var dataForNotification vendors.NotificationMessage
	dataForNotification.Title = data.Title
	dataForNotification.Body = data.Body
	dataForNotification.Data = vendors.JsonDataType(data.JsonData)

	vendors.SendNotification(tokens, dataForNotification)

	c.JSON(200, gin.H{
		"message": "success",
	})
}
