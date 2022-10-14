package app

import (
	"employeeSelfService/config"
	"employeeSelfService/database"
	handlerAbsenConfiguration "employeeSelfService/handler/absenConfiguration"
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

	// customerRepository := repostiory.NewCustomerRepository(dbClient)
	// customerService := service.NewCustomerService(customerRepository)
	// customerHandler := CustomerHandler{customerService}

	// // prepare handle auth login
	// userRepo := repostiory.NewUserRepository(dbClient)
	// authService := service.NewAuthService(userRepo)
	// authHandler := AuthHandler{authService}

	// // setup data admin
	// userRepo.SetupAdminDummy()

	// // prepare handle products
	// productRepo := repostiory.NewProductRepository(dbClient)
	// productService := service.NewProductService(productRepo)
	// productHandler := ProductHandler{productService}

	// // setup dummy product
	// productRepo.SetupProductDummy()

	r := gin.Default()

	// // productRoute := r.Group("/products")
	// // productRoute.Use(middleware.AuthMiddleware)
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

	// r.POST("/login", authHandler.LoginHandler)

	// r.POST("/products", middleware.IsAdminMiddleware(), productHandler.SaveProductHandler)
	// r.DELETE("/products/:productId", middleware.IsAdminMiddleware(), productHandler.DeleteProductHandler)
	// r.GET("/products", middleware.AuthMiddleware(), productHandler.GetAlProductHandler)
	// r.GET("/products/:productId", middleware.AuthMiddleware(), productHandler.GetProdutById)
	// r.PUT("/products/:productId", middleware.IsAdminMiddleware(), productHandler.UpdateProductHandler)

	// give info where server and port app running
	logger.Info(fmt.Sprintf("start server on  %s:%s ...", config.SERVER_ADDRESS, config.SERVER_PORT))

	// run server
	r.Run(fmt.Sprintf("%s:%s", config.SERVER_ADDRESS, config.SERVER_PORT))

}
