package app

import (
	"employeeSelfService/config"
	"employeeSelfService/database"
	handlerAbsenConfiguration "employeeSelfService/handler/absenConfiguration"
	handlerClient "employeeSelfService/handler/client"
	handlerDashboard "employeeSelfService/handler/dashboard"
	handlerEmailValidation "employeeSelfService/handler/emailValidation"
	handlerLogin "employeeSelfService/handler/login"
	handlerProject "employeeSelfService/handler/project"
	handlerRegister "employeeSelfService/handler/register"
	"employeeSelfService/helper"
	"employeeSelfService/logger"
	"employeeSelfService/middleware"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func Start() {
	// loading env variabel
	config.SetupEnv(".env")

	// check all variables are loaded
	config.SanityCheck()
	dbClient := database.GetClientDb()

	// check app run for develope or production
	if os.Getenv("TESTING") == "true" {
		// panic("this is end point")
		helper.TruncateAllTable(dbClient)
		database.SetupDataDummyTest(dbClient)
	}

	// prepare handler
	registerHandler := handlerRegister.NewHandlerRegister(dbClient)
	loginHandler := handlerLogin.NewHandlerLogin(dbClient)
	emailValidationHandler := handlerEmailValidation.NewHandlerEmailValidation(dbClient)
	absenConfigurationHandler := handlerAbsenConfiguration.NewHandlerAbsenConfiguration(dbClient)
	dashboardHandler := handlerDashboard.NewHandlerDashboard(dbClient)
	clientHandler := handlerClient.NewHandlerClient(dbClient)
	projectHandler := handlerProject.NewHandlerProject(dbClient)

	// setup server gin
	r := gin.Default()

	r.POST("/register", registerHandler.RegisterHandler)
	r.POST("/login", loginHandler.LoginHandler)

	r.GET("/dashboard/:user_id", dashboardHandler.GetDashboardHandler)
	isAdmin := r.Group("/")

	// is admin route middleware
	isAdmin.Use(middleware.IsAdmin(dbClient))
	{
		// route for konfigurasi or PIC
		konfigurasi := isAdmin.Group("/konfigurasi")
		{
			konfigurasi.POST("/:user_id/email", emailValidationHandler.SaveEmailValidation)
			konfigurasi.POST("/:user_id/kehadiran", absenConfigurationHandler.SaveAbsenConfiguration)
		}

		// route for client
		client := isAdmin.Group("/client")
		{
			client.POST("/client/:user_id", clientHandler.SaveClient)
			client.PUT("/client/:user_id", clientHandler.UpdateClient)
			client.DELETE("/client/:user_id/:client_id", clientHandler.DeleteClient)
		}

		// route for project
		project := isAdmin.Group("/project/:user_id")
		{
			project.POST("", projectHandler.SaveProject)
			project.PUT("", projectHandler.UpdateProject)
			project.DELETE("/:project_id", projectHandler.DeleteProject)
		}
	}

	// route for client for not admin
	client := r.Group("/client")
	{
		client.GET("", clientHandler.GetAllClient)
		client.GET("/:client_id", clientHandler.GetClientById)
	}

	// route for project not admin
	project := r.Group("/project")
	{
		project.GET("", projectHandler.GetAllProject)
		project.GET("/:project_id", projectHandler.GetProjectById)
	}

	// give info where server and port app running
	logger.Info(fmt.Sprintf("start server on  %s:%s ...", config.SERVER_ADDRESS, config.SERVER_PORT))

	// run server
	r.Run(fmt.Sprintf("%s:%s", config.SERVER_ADDRESS, config.SERVER_PORT))

}
