// Package controllers ...
package controllers

import (
	"server/config"
	"server/models"

	"github.com/gin-gonic/gin"
)

// IndexAssets ..
func IndexAssets(c *gin.Context) {
	var services []models.Services
	config.DB.Find(&services)

	c.JSON(200, gin.H{
		"services": services,
	})
}
