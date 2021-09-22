// Package app ...
package app

import (
	"server/config"
	"server/routes"
	"server/vendors"
)

// Installing ..
func Installing() {
	// Setup The DB
	config.SetupDB()

	config.ServerInformations()

	// SetupPassport
	vendors.SetupPassport()
	vendors.CreateAdmin()

	// setup routes
	routes.Setup()
}
