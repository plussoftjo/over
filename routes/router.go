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

	// --------- DriverRegister ------- //
	driverAuth := r.Group("/driverAuth")
	driverAuth.POST("/register", controllers.RegisterDriver)
	driverAuth.POST("/changeRegisterStatus", controllers.ChangeRegisterStatus)
	driverAuth.POST("/createValue", controllers.CreateDriverValues)
	driverAuth.POST("/createDriverCar", controllers.CreateDriverCar)

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
	driver.POST("/nearbyMainDrivers", controllers.GetMainNearbyDrivers)
	driver.GET("/getDriverRate/:id", controllers.GetDriverRate)
	driver.POST("/changeDriverStatus", controllers.ChangeDriverStatus)

	// --------------- Driver Controller ----------- //
	rider := r.Group("/rider")
	rider.POST("/storeRiderLocation", controllers.StoreRiderLocation)

	// ---------- Countries & Cites & Areas ------------ //
	countries := r.Group("/countries")
	countries.POST("/store", controllers.StoreCountry)
	countries.GET("/index", controllers.IndexCountries)
	countries.GET("/destroy/:id", controllers.DestroyCountry)
	countries.POST("/update", controllers.UpdateCountry)

	// ---------- PromosCode ------------ //
	promosCodes := r.Group("/promoCodes")
	promosCodes.POST("/store", controllers.StorePromoCode)
	promosCodes.GET("/index", controllers.IndexPromoCodes)
	promosCodes.GET("/destroy/:id", controllers.DestroyPromoCode)
	promosCodes.POST("/update", controllers.UpdatePromoCode)
	promosCodes.POST("/checkPromoCode", controllers.CheckPromoCode)

	// ---------- Services ------------ //
	services := r.Group("/services")
	services.POST("/store", controllers.StoreService)
	services.GET("/index", controllers.IndexServices)
	services.GET("/destroy/:id", controllers.DestroyService)
	services.POST("/update", controllers.UpdateService)

	// --------------- Application Controller ----------- //
	application := r.Group("/application")
	application.GET("/indexAssets", controllers.IndexAssets)
	application.POST("/storeNotificationToken", controllers.StoreNotificationToken)
	application.POST("/sendNotification", controllers.SendNotification)

	wallet := r.Group("/wallet")
	wallet.POST("/updateWallet", controllers.UpdateWallet)
	wallet.GET("/indexWalletLogs/:id", controllers.IndexWalletLogs)

	// Booking Chat controller
	bookingChat := r.Group("/bookingChat")
	bookingChat.POST("/store", controllers.StoreBookingChat)
	bookingChat.GET("/index/:id", controllers.IndexBookingChat)

	// BlockController
	blocking := r.Group("/blocking")
	blocking.POST("/store", controllers.StoreBlock)

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
	booking.GET("/history/:id", controllers.IndexUserHistory)
	booking.GET("/showBooking/:id", controllers.ShowBooking)

	// UserStatus ..
	userStatus := r.Group("/userStatus")
	userStatus.POST("/store", controllers.StoreUserStatus)

	// --------------- Dashboard Controller ----------- //

	// --------------- Drivers Controller ----------- //
	driversDashboard := r.Group("/driversDashboard")
	driversDashboard.GET("/indexDrivers/:type", controllers.IndexDrivers)
	driversDashboard.POST("/toggleUserBlock", controllers.ToggleUserBlock)
	driversDashboard.GET("/showDriver/:id", controllers.ShowDriver)
	driversDashboard.GET("/approveDriverRegister/:id", controllers.ApproveDriverRegister)
	driversDashboard.GET("/cancelDriverRegister/:id", controllers.CancelDriverRegister)

	// --------------- Drivers Controller ----------- //
	bookingsDashboard := r.Group("/bookingDashboard")
	bookingsDashboard.GET("/IndexBooking/:type", controllers.IndexBooking)
	bookingsDashboard.GET("/showBooking/:id", controllers.ShowBookingDashboard)

	// Dashboard
	dashboard := r.Group("/dashboard")
	dashboard.GET("/indexAllClients", controllers.IndexAllClients)
	dashboard.GET("/showUser/:id", controllers.ShowUser)

	r.GET("/ws/driver/:id", func(c *gin.Context) {
		id := c.Param("id")
		serveWs(c.Writer, c.Request, id)
	})

	r.Run(":8082")
}
