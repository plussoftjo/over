// Package controllers ...
package controllers

import (
	"net/http"
	"server/config"
	"server/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// StoreBooking ..
func StoreBooking(c *gin.Context) {

	//Error Codes
	// 101 => bad in store

	var bookingData models.Booking
	if err := c.ShouldBindJSON(&bookingData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Before set new booking
	// check if has notification before and remove it and set new one
	var lastBooking int64
	config.DB.Model(&models.Booking{}).Where("user_id = ?", bookingData.UserID).Where("status NOT IN (5,6)").Count(&lastBooking)
	if lastBooking != 0 {
		var lastBookingID int
		config.DB.Model(&models.Booking{}).Where("user_id = ?", bookingData.UserID).Where("status NOT IN (5,6)").Pluck("id", &lastBookingID)
		config.DB.Unscoped().Delete(&models.Booking{}, lastBookingID)
	}

	err := config.DB.Create(&bookingData).Error
	if err != nil {
		// There are error
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"code":  101,
		})
		return
	}

	c.JSON(200, gin.H{
		"booking": bookingData,
	})
}

// UpdateBooking ..
func UpdateBooking(c *gin.Context) {

	// UpdateBookingType ..
	type UpdateBookingType struct {
		BookingID uint `json:"bookingID"`
		Status    int  `json:"status"`
		DriverID  uint `json:"driverID"`
	}

	// Codes Types
	// 100 => bind Error
	// 101 => Get record Error
	// 102

	var bookingUpdateData UpdateBookingType

	// Bind JSON
	if err := c.ShouldBindJSON(&bookingUpdateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": 100})
		return
	}

	var bookingData models.Booking
	err := config.DB.Where("id = ?", bookingUpdateData.BookingID).First(&bookingData).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": 101})
		return
	}

	if bookingUpdateData.Status == 1 {
		bookingData.DriverID = bookingUpdateData.DriverID
	}
	bookingData.Status = bookingUpdateData.Status

	SaveError := config.DB.Save(&bookingData).Error
	if SaveError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": 102})
		return
	}

	c.JSON(200, gin.H{
		"message": "complete",
	})

}

// CheckIfUserHaveOrder ..
func CheckIfUserHaveOrder(c *gin.Context) {
	ID := c.Param("id")
	var ordersCount int
	err := config.DB.Model(&models.Booking{}).Where("user_id = ?", ID).Where("status NOT IN (5,6)").Count(&ordersCount).Error
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
			"code":  100,
		})
		return
	}
	if ordersCount == 0 {
		c.JSON(200, gin.H{
			"hasOrder": false,
		})
		return
	} else {
		var booking models.Booking
		recordError := config.DB.Where("user_id = ?", ID).Where("status NOT IN (5,6)").Preload("Service").Preload("Driver", func(db *gorm.DB) *gorm.DB {
			return db.Scopes(models.DriverWithDetails)
		}).First(&booking).Error
		if recordError != nil {
			c.JSON(500, gin.H{
				"code":  101,
				"error": recordError.Error(),
			})
		}

		if booking.Status == 0 && booking.BookingType == "now" {
			config.DB.Unscoped().Delete(&models.Booking{}, booking.ID)
			c.JSON(200, gin.H{
				"hasOrder": false,
			})
			return
		}
		c.JSON(200, gin.H{
			"hasOrder": true,
			"booking":  booking,
		})
		return
	}
}

// CheckIfDriverHaveOrder ..
func CheckIfDriverHaveOrder(c *gin.Context) {
	ID := c.Param("id")
	var ordersCount int
	err := config.DB.Model(&models.Booking{}).Where("driver_id = ?", ID).Where("status NOT IN (5,6)").Count(&ordersCount).Error
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
			"code":  100,
		})
		return
	}
	if ordersCount == 0 {
		c.JSON(200, gin.H{
			"hasOrder": false,
		})
		return
	} else {
		var booking models.Booking
		recordError := config.DB.Where("driver_id = ?", ID).Where("status NOT IN (5,6)").Preload("Service").Preload("Driver", func(db *gorm.DB) *gorm.DB {
			return db.Scopes(models.DriverWithDetails)
		}).Preload("User").First(&booking).Error
		if recordError != nil {
			c.JSON(500, gin.H{
				"code":  101,
				"error": recordError.Error(),
			})
		}
		c.JSON(200, gin.H{
			"hasOrder": true,
			"booking":  booking,
		})
		return
	}
}

// ShowBooking ..
func ShowBooking(c *gin.Context) {
	ID := c.Param("id")
	var data models.Booking

	err := config.DB.Preload("Service").Preload("Driver", func(db *gorm.DB) *gorm.DB {
		return db.Scopes(models.DriverWithDetails)
	}).Preload("User").Where("id = ?", ID).First(&data).Error
	if err != nil {
		c.JSON(500, gin.H{
			"err":  err.Error(),
			"code": 101,
		})
		return
	}

	c.JSON(200, data)
}

// FinishBookingAndRateFromUser ..
func FinishBookingAndRateFromUser(c *gin.Context) {
	type DataType struct {
		OrderID  uint            `json:"orderID"`
		UserRate models.UserRate `json:"userRate"`
	}

	var data DataType
	c.ShouldBindJSON(&data)

	// Change the booking Status
	config.DB.Model(&models.Booking{}).Where("id = ?", data.OrderID).Updates(&models.Booking{
		Status: 5,
	})

	config.DB.Create(&data.UserRate)

	c.JSON(200, gin.H{
		"message": "success",
	})
}

// Steps

// OnDriverArrive ..
func OnDriverArrive(c *gin.Context) {
	// OnDriverArriveType ..
	type OnDriverArriveType struct {
		BookingID  uint      `json:"bookingID"`
		ArriveDate time.Time `json:"arriveDate"`
	}

	var data OnDriverArriveType
	c.ShouldBindJSON(&data)

	// Update
	config.DB.Model(&models.Booking{}).Where("id = ?", data.BookingID).Updates(&models.Booking{
		Status:     2,
		ArriveDate: data.ArriveDate,
	})
}

// OnStartTrip ..
func OnStartTrip(c *gin.Context) {
	type OnStartTripType struct {
		BookingID     uint      `json:"bookingID"`
		StartTripDate time.Time `json:"startTripDate"`
		WaitingTime   float64   `json:"waitingTime"`
	}

	var data OnStartTripType
	c.ShouldBindJSON(&data)

	err := config.DB.Model(&models.Booking{}).Where("id = ?", data.BookingID).Updates(&models.Booking{
		Status:        3,
		WaitingTime:   data.WaitingTime,
		StartTripDate: data.StartTripDate,
	}).Error
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
}

// UpdateMetersInBooking ..
func UpdateMetersInBooking(c *gin.Context) {
	type DataType struct {
		BookingID uint `json:"bookingID"`
		Meters    int  `json:"meters"`
	}

	var data DataType
	c.ShouldBindJSON(&data)

	err := config.DB.Model(&models.Booking{}).Where("id = ?", data.BookingID).Updates(&models.Booking{
		Meters: data.Meters,
	}).Error

	if err != nil {
		c.JSON(500, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"message": "success"})
}

// EndTrip ..
func EndTrip(c *gin.Context) {
	type DataType struct {
		BookingID   uint    `json:"bookingID"`
		TripTime    float64 `json:"tripTime"`
		TripDetails string  `json:"tripDetails"`
	}

	var data DataType
	c.ShouldBindJSON(&data)

	err := config.DB.Model(&models.Booking{}).Where("id = ?", data.BookingID).Updates(&models.Booking{
		TripTime:    data.TripTime,
		TripDetails: data.TripDetails,
		DriverRate:  true,
		Status:      4,
	}).Error
	if err != nil {
		c.JSON(500, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
	})

}

// IndexUserHistory ..
func IndexUserHistory(c *gin.Context) {
	ID := c.Param("id")
	var history []models.Booking

	err := config.DB.Where("user_id = ?", ID).Find(&history).Error
	if err != nil {
		c.JSON(500, gin.H{
			"err":  err.Error(),
			"code": 101,
		})
		return
	}

	c.JSON(200, history)

}
