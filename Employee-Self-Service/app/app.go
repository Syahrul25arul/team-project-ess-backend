package app

import (
	"employeeSelfService/config"
	"employeeSelfService/database"
	handlerAbsenConfiguration "employeeSelfService/handler/absenConfiguration"
	handlerClient "employeeSelfService/handler/client"
	handlerDashboard "employeeSelfService/handler/dashboard"
	handlerEmailValidation "employeeSelfService/handler/emailValidation"
	handlerLogin "employeeSelfService/handler/login"
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
	clientHandler := handlerClient.NewHandlerClient(dbClient)

	// setup server gin
	r := gin.Default()

	r.POST("/register", registerHandler.RegisterHandler)
	r.POST("/login", loginHandler.LoginHandler)

	r.GET("/dashboard/:user_id", dashboardHandler.GetDashboardHandler)
	isAdmin := r.Group("/konfigurasi")

	// is admin route middleware
	isAdmin.Use(middleware.IsAdmin(dbClient))
	{
		isAdmin.POST("/:user_id/email", emailValidationHandler.SaveEmailValidation)
		isAdmin.POST("/:user_id/kehadiran", absenConfigurationHandler.SaveAbsenConfiguration)
		isAdmin.GET("/client/:user_id", clientHandler.GetAllClient)
		isAdmin.GET("/client/:user_id/:client_id", clientHandler.GetClientById)
		isAdmin.POST("/client/:user_id", clientHandler.SaveClient)
		isAdmin.PUT("/client/:user_id", clientHandler.UpdateClient)
		isAdmin.DELETE("/client/:user_id/:client_id", clientHandler.DeleteClient)
	}

	// give info where server and port app running
	logger.Info(fmt.Sprintf("start server on  %s:%s ...", config.SERVER_ADDRESS, config.SERVER_PORT))

	// run server
	r.Run(fmt.Sprintf("%s:%s", config.SERVER_ADDRESS, config.SERVER_PORT))

}
