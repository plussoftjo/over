// Package models ..
package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Booking ..
type Booking struct {
	gorm.Model
	UserID          uint      `json:"userID"`
	Origin          string    `json:"origin"`
	Destination     string    `json:"destination"`
	DestinationType string    `json:"destinationType" gorm:"default:'set'"` // set => set the destination , open => open destination for ride now
	Status          int       `json:"status" gorm:"default:0"`              // 0 => NewOrder , 1 => TakenDriverInTheWay, 2 => DriverArrive, 3 => OnTrip, 4 => EndingNotRated, 5 => EndingWithRated, 6 => Canceled
	Note            string    `json:"string"`
	DriverID        uint      `json:"driverID"`
	TripDetails     string    `json:"tripDetails"`
	BookingType     string    `json:"bookingType" gorm:"default:'now'"`
	BookingTime     time.Time `json:"bookingTime"`
	ServiceID       uint      `json:"serviceID"`
	StartTripDate   time.Time `json:"startTripDate"`
	Meters          int       `json:"meters" gorm:"default:0"`
	TripTime        float64   `json:"tripTime" gorm:"default:0"`
	ArriveDate      time.Time `json:"arriveDate"`
	WaitingTime     float64   `json:"waitingTime"` // In Minute
	DriverRate      bool      `json:"driverRate"`
	Driver          User      `json:"driver" gorm:"foreignKey:DriverID;references:ID"`
	User            User      `json:"rider" gorm:"foreignKey:UserID;references:ID"`
	Service         Services  `json:"service" gorm:"foreignKey:ServiceID;references:ID"`
}
