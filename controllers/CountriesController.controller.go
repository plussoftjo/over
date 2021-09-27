// Package controllers ...
package controllers

import (
	"server/config"
	"server/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

// ------------- Country -------------//

// StoreCountry ..
func StoreCountry(c *gin.Context) {
	// Register var and bind json
	var country models.Countries
	if err := c.ShouldBindJSON(&country); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// StoreInDB
	if err := config.DB.Create(&country).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, country)
}

// IndexCountries ..
func IndexCountries(c *gin.Context) {
	// Register countries in var
	var countries []models.Countries
	config.DB.Find(&countries)
	c.JSON(200, countries)
}

// DestroyCountry ..
func DestroyCountry(c *gin.Context) {
	ID := c.Param("id")
	config.DB.Delete(&models.Countries{}, ID)
	// Return Countries list
	var countries []models.Countries
	config.DB.Find(&countries)
	c.JSON(200, countries)
}

// UpdateCountry ..
func UpdateCountry(c *gin.Context) {
	var country models.Countries
	if err := c.ShouldBindJSON(&country); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Model(&country).Update(&country).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var countries []models.Countries
	config.DB.Find(&countries)
	c.JSON(200, gin.H{
		"country":   country,
		"countries": countries,
	})
}

// ------------- End Country ------------ //
