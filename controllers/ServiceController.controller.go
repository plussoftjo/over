// Package controllers ...
package controllers

import (
	"server/config"
	"server/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

// ------------- Services -------------//

// StoreService ..
func StoreService(c *gin.Context) {
	var service models.Services
	if err := c.ShouldBindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// StoreInDB
	if err := config.DB.Create(&service).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, service)
}

// IndexServices ..
func IndexServices(c *gin.Context) {
	var services []models.Services
	config.DB.Find(&services)
	c.JSON(200, services)
}

// DestroyService ..
func DestroyService(c *gin.Context) {
	ID := c.Param("id")
	config.DB.Delete(&models.Services{}, ID)
	var services []models.Services
	config.DB.Find(&services)
	c.JSON(200, services)
}

// UpdateService ..
func UpdateService(c *gin.Context) {
	var service models.Services
	if err := c.ShouldBindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Model(&service).Update(&service).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var services []models.Services
	config.DB.Find(&services)
	c.JSON(200, gin.H{
		"service":  service,
		"services": services,
	})
}

// ------------- End Services ------------ //
