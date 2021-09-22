package controllers

import (
	"net/http"
	"server/config"
	"server/models"

	"github.com/gin-gonic/gin"
)

// UpdateDriverDetails ..
func UpdateDriverDetails(c *gin.Context) {
	var driverDetails models.DriverDetails
	if err := c.ShouldBindJSON(&driverDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var newDriverDetails models.DriverDetails
	err := config.DB.Where("user_id = ?", driverDetails.UserID).First(&newDriverDetails).Error
	if err != nil {

		config.DB.Create(&driverDetails)
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	newDriverDetails.Latitude = driverDetails.Latitude
	newDriverDetails.Longitude = driverDetails.Longitude
	newDriverDetails.Heading = driverDetails.Heading
	newDriverDetails.UserID = driverDetails.UserID

	config.DB.Save(&newDriverDetails)

	c.JSON(200, gin.H{
		"message": "Success",
	})
}

// GetNearbyDrivers ..
func GetNearbyDrivers(c *gin.Context) {
	// LatLng
	type LatLng struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}

	type Result struct {
		UserID    uint    `json:"userID"`
		ID        uint    `json:"ID"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
		Distance  float64 `json:"distance"`
		Heading   float64 `json:"heading"`
		Status    int64   `json:"status"`
	}

	var latLng LatLng
	c.ShouldBindJSON(&latLng)

	var result []Result
	// Select Latitude And Longitude
	config.DB.Raw(`SELECT latitude, longitude,user_id,id,status,heading, SQRT(
		POW(69.1 * (latitude - ?), 2) +
		POW(69.1 * (? - longitude) * COS(latitude / 57.3), 2)) AS distance
	FROM driver_details HAVING distance < 200 ORDER BY distance;`, latLng.Latitude, latLng.Longitude).Scan(&result)

	c.JSON(200, gin.H{
		"result": result,
	})
}

// GetDriverRate ..
func GetDriverRate(c *gin.Context) {
	ID := c.Param("id")

	var userRate []models.UserRate
	err := config.DB.Where("user_id = ?", ID).Find(&userRate).Error
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	var rateNumber float64
	rateNumber = 4.3

	// Check if have rate or not
	if len(userRate) != 0 {
		sumTheRate := 0.0
		for _, r := range userRate {
			sumTheRate = sumTheRate + r.Stars
		}
		rateNumber = sumTheRate / float64(len(userRate))
	}

	c.JSON(200, gin.H{
		"rate": rateNumber,
	})

}

// ChangeDriverStatus ..
func ChangeDriverStatus(c *gin.Context) {
	type DataType struct {
		UserID uint  `json:"userID"`
		Status int64 `json:"status"`
	}

	var data DataType
	c.ShouldBindJSON(&data)

	// Fetch the driver details
	err := config.DB.Model(&models.DriverDetails{}).Where("user_id = ?", data.UserID).Updates(models.DriverDetails{
		Status: data.Status,
	}).Error

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}
