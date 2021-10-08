// Package controllers ...
package controllers

import (
	"server/config"
	"server/models"
	"server/vendors"

	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// RegisterDriver ..
func RegisterDriver(c *gin.Context) {
	// Get user data
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// HashThePassword
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// SetThePasswordToTheUser
	user.Password = string(hashedPassword)
	// CreateTheUser
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": 101})
		return
	}
	//CreateTheToken
	token, err := vendors.CreateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// NowTheSteps

	// FirstListCreateWallet
	config.DB.Create(&models.Wallet{
		UserID:  user.ID,
		Balance: 0,
	})

	// CreateDriverDetails ..
	config.DB.Create(&models.DriverDetails{UserID: user.ID, Status: 1})

	c.JSON(http.StatusOK, gin.H{"user": user, "token": token})
}

// ChangeRegisterStatus ..
func ChangeRegisterStatus(c *gin.Context) {
	type ChangeRegisterStatusType struct {
		UserID         uint `json:"userID"`
		RegisterStatus int  `json:"registerStatus"`
	}

	var data ChangeRegisterStatusType
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := config.DB.Model(&models.User{}).Where("id = ?", data.UserID).Updates(&models.User{RegisterStatus: data.RegisterStatus}).Error
	if err != nil {
		c.JSON(500, gin.H{
			"err":  err.Error(),
			"code": 101,
		})
		return
	}

	c.JSON(200, gin.H{"message": "success"})
}

// CreateDriverValues ..
func CreateDriverValues(c *gin.Context) {
	var values []models.DriverValues
	c.ShouldBindJSON(&values)

	for _, value := range values {
		config.DB.Create(&value)
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
}

// CreateDriverCar ..
func CreateDriverCar(c *gin.Context) {
	var car models.DriversCar
	c.ShouldBindJSON(&car)

	config.DB.Create(&car)

	c.JSON(200, gin.H{
		"message": "success",
	})
}
