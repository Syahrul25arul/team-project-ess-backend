package app

import (
	"employeeSelfService/config"
	"employeeSelfService/database"
	handlerAbsenConfiguration "employeeSelfService/handler/absenConfiguration"
	handlerClient "employeeSelfService/handler/client"
	handlerDashboard "employeeSelfService/handler/dashboard"
	handlerEmailValidation "employeeSelfService/handler/emailValidation"
	handlerLogin "employeeSelfService/handler/login"
	handlerPosition "employeeSelfService/handler/position"
	handlerRegister "employeeSelfService/handler/register"
	"employeeSelfService/logger"
	"employeeSelfService/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Start() {
	// loading env variabel
	config.SetupEnv(".env")

	// check all variables are loaded
	config.SanityCheck()
	dbClient := database.GetClientDb()

	// prepare handler
	registerHandler := handlerRegister.NewHandlerRegister(dbClient)
	loginHandler := handlerLogin.NewHandlerLogin(dbClient)
	emailValidationHandler := handlerEmailValidation.NewHandlerEmailValidation(dbClient)
	absenConfigurationHandler := handlerAbsenConfiguration.NewHandlerAbsenConfiguration(dbClient)
	dashboardHandler := handlerDashboard.NewHandlerDashboard(dbClient)
	positionHandler := handlerPosition.NewHandlerPosition(dbClient)
	clientHandler := handlerClient.NewHandlerClient(dbClient)

	// setup server gin
	r := gin.Default()

	r.POST("/register", registerHandler.RegisterHandler)
	r.POST("/login", loginHandler.LoginHandler)
	r.GET("/dashboard/:user_id", dashboardHandler.GetDashboardHandler)

	// isAdmin := r.Group("/konfigurasi")
	isAdmin := r.Group("/")

	// is admin route middleware
	isAdmin.Use(middleware.IsAdmin(dbClient))
	{
		konfigurasi := isAdmin.Group("/konfigurasi")
		konfigurasi.Use()
		{
			konfigurasi.POST("/:user_id/email", emailValidationHandler.SaveEmailValidation)
			konfigurasi.POST("/:user_id/kehadiran", absenConfigurationHandler.SaveAbsenConfiguration)
			konfigurasi.GET("/client/:user_id", clientHandler.GetAllClient)
			konfigurasi.GET("/client/:user_id/:client_id", clientHandler.GetClientById)
			konfigurasi.POST("/client/:user_id", clientHandler.SaveClient)
			konfigurasi.PUT("/client/:user_id", clientHandler.UpdateClient)
			konfigurasi.DELETE("/client/:user_id/:client_id", clientHandler.DeleteClient)
		}

		position := isAdmin.Group("/position")
		position.Use()
		{
			position.POST("/", positionHandler.SavePosition)
			position.GET("/:id_position", positionHandler.GetPositionById)
			position.DELETE("/:id_position", positionHandler.DeletePosition)
			position.PUT("/:id_position", positionHandler.UpdatePosition)
		}

	}

	// give info where server and port app running
	logger.Info(fmt.Sprintf("start server on  %s:%s ...", config.SERVER_ADDRESS, config.SERVER_PORT))

	// run server
	r.Run(fmt.Sprintf("%s:%s", config.SERVER_ADDRESS, config.SERVER_PORT))

}
