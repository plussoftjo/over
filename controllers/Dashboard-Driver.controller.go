package controllers

import (
	"server/config"
	"server/models"

	"github.com/gin-gonic/gin"
)

// IndexDrivers ...
func IndexDrivers(c *gin.Context) {
	Type := c.Param("type")

	var drivers []models.User
	if Type == "actives" {
		err := config.DB.Where("user_type = ?", 2).Where("register_status = ?", 101).Scopes(models.DriverWithDashboardDetails).Find(&drivers).Error
		if err != nil {
			c.JSON(500, gin.H{
				"err":  err.Error(),
				"code": 101,
			})
			return
		}
	}

	if Type == "driverRequests" {
		err := config.DB.Where("user_type = ?", 2).Where("register_status = ?", 100).Scopes(models.DriverWithDashboardDetails).Find(&drivers).Error
		if err != nil {
			c.JSON(500, gin.H{
				"err":  err.Error(),
				"code": 101,
			})
			return
		}
	}

	c.JSON(200, drivers)
}

// ToggleUserBlock ..
func ToggleUserBlock(c *gin.Context) {
	type ToggleUserBlockType struct {
		UserID uint `json:"userID"`
		Block  int  `json:"block"`
	}

	var data ToggleUserBlockType
	c.ShouldBindJSON(&data)

	var user models.User
	err := config.DB.Where("id = ?", data.UserID).First(&user).Error
	if err != nil {
		c.JSON(500, gin.H{
			"err":  err.Error(),
			"code": 101,
		})
		return
	}

	user.Block = data.Block
	config.DB.Save(&user)

	c.JSON(200, gin.H{
		"message": "success",
	})

}

// ShowDriver ..
func ShowDriver(c *gin.Context) {
	ID := c.Param("id")
	var user models.User

	err := config.DB.Where("id = ?", ID).Scopes(models.DriverWithDashboardDetails).First(&user).Error
	if err != nil {
		c.JSON(500, gin.H{
			"err":  err.Error(),
			"code": 101,
		})
		return
	}

	var bookings []models.Booking
	config.DB.Where("driver_id = ?", ID).Scopes(models.BookingWithDetails).Find(&bookings)

	c.JSON(200, gin.H{
		"user":     user,
		"bookings": bookings,
	})
}

// ApproveDriverRegister ..
func ApproveDriverRegister(c *gin.Context) {
	ID := c.Param("id")
	var user models.User

	err := config.DB.Where("id = ?", ID).First(&user).Error
	if err != nil {
		c.JSON(500, gin.H{
			"err":  err.Error(),
			"code": 101,
		})
		return
	}

	user.RegisterStatus = 101

	config.DB.Save(&user)

	c.JSON(200, gin.H{
		"message": "success",
	})
}

// CancelDriverRegister ..
func CancelDriverRegister(c *gin.Context) {
	ID := c.Param("id")
	var user models.User

	err := config.DB.Where("id = ?", ID).First(&user).Error
	if err != nil {
		c.JSON(500, gin.H{
			"err":  err.Error(),
			"code": 101,
		})
		return
	}

	user.RegisterStatus = 500

	config.DB.Save(&user)

	c.JSON(200, gin.H{
		"message": "success",
	})
}
