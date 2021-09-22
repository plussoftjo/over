// Package controllers ..
package controllers

import (
	"image"
	"os"
	"path/filepath"
	"server/config"
	"server/vendors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// UpdateImage ..
func UpdateImage(c *gin.Context) {
	file, _ := c.FormFile("image")
	imageType := c.Param("imageType")

	filename := filepath.Base(file.Filename)
	id := uuid.NewString()

	if err := c.SaveUploadedFile(file, config.ServerInfo.PublicPath+"public/"+imageType+"/"+filename); err != nil {
		c.JSON(500, gin.H{
			"error":   err.Error(),
			"message": "error",
		})
		return
	}

	afterUploadImage, _ := os.Open(config.ServerInfo.PublicPath + "public/" + imageType + "/" + filename)
	_, format, _ := image.DecodeConfig(afterUploadImage)
	afterUploadImage.Close()

	id = id + "." + format
	vendors.ResizeImage(filename, id, config.ServerInfo.PublicPath+"public/"+imageType+"/", format)
	c.JSON(200, gin.H{
		"image":   id,
		"message": "success",
	})
}
