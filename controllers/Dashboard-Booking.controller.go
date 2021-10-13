package controllers

import (
	"server/config"
	"server/models"

	"github.com/gin-gonic/gin"
)

// IndexBooking ..
func IndexBooking(c *gin.Context) {
	QueryType := c.Param("type")

	var bookings []models.Booking

	if QueryType == "newOrders" {
		err := config.DB.Where("status = ?", 0).Find(&bookings).Error
		if err != nil {
			c.JSON(500, gin.H{
				"err":  err.Error(),
				"code": 101,
			})
			return
		}
	}
	if QueryType == "waitingDrivers" {
		err := config.DB.Where("status = ?", 1).Find(&bookings).Error
		if err != nil {
			c.JSON(500, gin.H{
				"err":  err.Error(),
				"code": 101,
			})
			return
		}
	}
	if QueryType == "waitingRiders" {
		err := config.DB.Where("status = ?", 2).Find(&bookings).Error
		if err != nil {
			c.JSON(500, gin.H{
				"err":  err.Error(),
				"code": 101,
			})
			return
		}
	}

	if QueryType == "onTrip" {
		err := config.DB.Where("status = ?", 3).Find(&bookings).Error
		if err != nil {
			c.JSON(500, gin.H{
				"err":  err.Error(),
				"code": 101,
			})
			return
		}
	}

	if QueryType == "waitingRate" {
		err := config.DB.Where("status = ?", 4).Find(&bookings).Error
		if err != nil {
			c.JSON(500, gin.H{
				"err":  err.Error(),
				"code": 101,
			})
			return
		}
	}

	if QueryType == "completedOrders" {
		err := config.DB.Where("status = ?", 5).Find(&bookings).Error
		if err != nil {
			c.JSON(500, gin.H{
				"err":  err.Error(),
				"code": 101,
			})
			return
		}
	}

	if QueryType == "cancelOrders" {
		err := config.DB.Where("status = ?", 6).Find(&bookings).Error
		if err != nil {
			c.JSON(500, gin.H{
				"err":  err.Error(),
				"code": 101,
			})
			return
		}
	}

	c.JSON(200, bookings)

}

// ShowBookingDashboard
func ShowBookingDashboard(c *gin.Context) {
	ID := c.Param("id")

	var booking models.Booking

	err := config.DB.Where("id = ?", ID).Scopes(models.BookingWithDetailsDashboard).First(&booking).Error
	if err != nil {
		c.JSON(500, gin.H{
			"err":  err.Error(),
			"code": 101,
		})
		return
	}

	c.JSON(200, booking)
}
