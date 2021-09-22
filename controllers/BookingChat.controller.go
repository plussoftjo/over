// Package controllers ...
package controllers

import (
	"server/config"
	"server/models"

	"github.com/gin-gonic/gin"
)

// StoreBookingChat ..
func StoreBookingChat(c *gin.Context) {
	var data models.BookingChats
	c.ShouldBindJSON(&data)

	err := config.DB.Create(&data).Error
	if err != nil {
		c.JSON(500, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(200, data)

}

// IndexBookingChat ..
func IndexBookingChat(c *gin.Context) {
	ID := c.Param("id")

	var messages []models.BookingChats
	config.DB.Where("booking_id = ?", ID).Order("id desc").Find(&messages)

	c.JSON(200, messages)
}
