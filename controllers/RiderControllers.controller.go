package controllers

import (
	"net/http"
	"server/config"
	"server/models"

	"github.com/gin-gonic/gin"
)

// StoreRiderLocation ..
func StoreRiderLocation(c *gin.Context) {
	var riderLocation models.RiderLocations

	if err := c.ShouldBindJSON(&riderLocation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&riderLocation)

	var riderLocationsList []models.RiderLocations

	if errCreate := config.DB.Where("user_id = ?", riderLocation.UserID).Find(&riderLocationsList).Error; errCreate != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errCreate.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message":            "Added Success",
		"riderLocation":      riderLocation,
		"riderLocationsList": riderLocationsList,
	})

}
