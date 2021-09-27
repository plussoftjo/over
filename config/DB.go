// Package config ...
package config

import (
	"server/models"

	"github.com/jinzhu/gorm"
	// Connect mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// models
)

// SetupDB ...

// DB ..
var DB *gorm.DB

// SetupDB ..
func SetupDB() {
	database, err := gorm.Open("mysql", "root:00962s00962S!@tcp(127.0.0.1:3306)/taxiapp?charset=utf8mb4&parseTime=True&loc=Local")

	// If Error in Connect
	if err != nil {
		panic(err)
	}
	// User Models Setup
	database.AutoMigrate(&models.User{})
	database.AutoMigrate(&models.AuthClients{})
	database.AutoMigrate(&models.AuthTokens{})
	database.AutoMigrate(&models.Roles{})
	database.AutoMigrate(&models.DriverDetails{})
	database.AutoMigrate(&models.RiderLocations{})
	database.AutoMigrate(&models.Services{})
	database.AutoMigrate(&models.Booking{})

	database.AutoMigrate(&models.UserRate{})
	database.AutoMigrate(&models.DriversCar{})
	database.AutoMigrate(&models.BookingChats{})

	database.AutoMigrate(&models.Wallet{})
	database.AutoMigrate(&models.WalletLogs{})

	database.AutoMigrate(&models.Countries{})
	database.AutoMigrate(&models.BlockList{})
	database.AutoMigrate(&models.UserStatus{})

	DB = database
}
