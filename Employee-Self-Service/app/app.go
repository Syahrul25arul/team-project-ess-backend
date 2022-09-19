package app

func Start() {
	// // loading env variabel
	// if err := godotenv.Load(); err != nil {
	// 	logger.Fatal("error loading file .env variable " + err.Error())
	// }

	// // check all variables are loaded
	// config.SanityCheck()
	// dbClient := database.GetClientDb()

	// // prepare handler customer
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

	// r := gin.Default()

	// // productRoute := r.Group("/products")
	// // productRoute.Use(middleware.AuthMiddleware)
	// r.POST("/register", customerHandler.RegisterCustomerHandler)
	// r.POST("/login", authHandler.LoginHandler)

	// r.POST("/products", middleware.IsAdminMiddleware(), productHandler.SaveProductHandler)
	// r.DELETE("/products/:productId", middleware.IsAdminMiddleware(), productHandler.DeleteProductHandler)
	// r.GET("/products", middleware.AuthMiddleware(), productHandler.GetAlProductHandler)
	// r.GET("/products/:productId", middleware.AuthMiddleware(), productHandler.GetProdutById)
	// r.PUT("/products/:productId", middleware.IsAdminMiddleware(), productHandler.UpdateProductHandler)

	// // give info where server and port app running
	// logger.Info(fmt.Sprintf("start server on  %s:%s ...", config.SERVER_ADDRESS, config.SERVER_PORT))

	// // run server
	// r.Run(fmt.Sprintf("%s:%s", config.SERVER_ADDRESS, config.SERVER_PORT))

}
