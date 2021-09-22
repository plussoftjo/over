// Package routes (Setup Routes Group)
package routes

import (
	"server/config"
	"server/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

// Setup >>>
func Setup() {
	r := gin.Default()
	go h.run()
	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token", "authorization", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowAllOrigins:  false,
		AllowOriginFunc:  func(origin string) bool { return true },
		MaxAge:           86400,
	}))
	// gin.SetMode(gin.ReleaseMode)
	r.Use(static.Serve("/public", static.LocalFile(config.ServerInfo.PublicPath+"public", true)))
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Success",
		})
	})

	// -------- Auth Groups ----------//

	// ~~~ Auth Group ~~~ //
	auth := r.Group("/auth")
	auth.POST("/login", controllers.LoginController)
	auth.POST("/register", controllers.RegisterController)
	auth.POST("/app/register", controllers.AppRegisterController)
	auth.POST("/driver/register", controllers.DriverRegisterController)
	auth.POST("/app/login", controllers.AppLoginController)
	auth.GET("/app/auth", controllers.AuthAppUser)
	auth.POST("/app/changePassword", controllers.ChangePassword)
	auth.GET("/auth", controllers.Auth)
	auth.GET("/users/index", controllers.UsersListIndex)
	auth.GET("/users/delete/:id", controllers.DeleteUser)
	auth.POST("/update", controllers.UpdateUser)
	auth.POST("/app/update", controllers.AppUpdateUser)
	auth.POST("/checkHasPhone", controllers.CheckIfHasPhone)
	auth.POST("/resetPassword", controllers.ResetPassword)

	// --------- Basics ------- //
	basics := r.Group("/basics")

	// UploadImage => For All
	basics.POST("/upload_image/:imageType", controllers.UpdateImage)

	// --------- User Controller ----------------- //
	user := r.Group("/users")
	// ~~~ User Roles ~~~ //
	user.POST("/roles/store", controllers.StoreUserRoles)
	user.POST("/roles/update", controllers.UpdateUserRole)
	user.GET("/roles/index", controllers.IndexUserRoles)
	user.GET("/roles/delete/:id", controllers.DeleteUserRole)
	// --------------- Employ Controller ----------- //
	user.POST("/employee/store", controllers.StoreEmployee)
	user.GET("/employee/index", controllers.IndexEmployee)
	user.GET("/employee/delete/:id", controllers.DeleteEmployee)
	user.POST("/employee/update", controllers.UpdateEmployee)

	// --------------- Driver Controller ----------- //
	driver := r.Group("/driver")
	driver.POST("/driverDetails/update", controllers.UpdateDriverDetails)
	driver.POST("/nearbyDrivers", controllers.GetNearbyDrivers)
	driver.GET("/getDriverRate/:id", controllers.GetDriverRate)
	driver.POST("/changeDriverStatus", controllers.ChangeDriverStatus)

	// --------------- Driver Controller ----------- //
	rider := r.Group("/rider")
	rider.POST("/storeRiderLocation", controllers.StoreRiderLocation)

	// --------------- Application Controller ----------- //
	application := r.Group("/application")
	application.GET("/indexAssets", controllers.IndexAssets)

	wallet := r.Group("/wallet")
	wallet.POST("/updateWallet", controllers.UpdateWallet)

	// Booking Chat controller
	bookingChat := r.Group("/bookingChat")
	bookingChat.POST("/store", controllers.StoreBookingChat)
	bookingChat.GET("/index/:id", controllers.IndexBookingChat)
	// ---------- Bookings Controller ------- //
	booking := r.Group("/booking")
	booking.POST("/storeBooking", controllers.StoreBooking)
	booking.POST("/updateBooking", controllers.UpdateBooking)
	booking.GET("/checkIfUserHaveOrder/:id", controllers.CheckIfUserHaveOrder)
	booking.GET("/CheckIfDriverHaveOrder/:id", controllers.CheckIfDriverHaveOrder)
	booking.POST("/finishBookingWithRateFromUser", controllers.FinishBookingAndRateFromUser)
	booking.POST("/onDriverArrive", controllers.OnDriverArrive)
	booking.POST("/onStartTrip", controllers.OnStartTrip)
	booking.POST("/updateMeters", controllers.UpdateMetersInBooking)
	booking.POST("/endTrip", controllers.EndTrip)

	r.GET("/ws/driver/:id", func(c *gin.Context) {
		id := c.Param("id")
		serveWs(c.Writer, c.Request, id)
	})

	r.Run(":8082")
}