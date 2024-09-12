package main

import (
	"Blog/dbschema"
	"Blog/internal/config"
	"Blog/internal/routes"
	"Blog/internal/server"
	"Blog/pkg/logger"
	"log"
)

const (
	defaultPort = "4200"
	logFilePath = "./logs"
	logFileName = "app_logs.db"
)

func main() {
	if err := logger.InitLogger(logFilePath, logFileName); err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}

	cfg, err := config.Load()
	if err != nil {
		logger.Fatal("Failed to load configuration: %v", err)
	}

	queries, blogDB := dbschema.ConnectCreateDoShitWithDB()
	defer blogDB.Close()

	appServices := routes.New(queries)
	appHandlers := appServices.CreateHandlers()

	server := server.New(cfg.Port, appHandlers)
	server.Run()
}

// func initConfig() {
// 	if err := godotenv.Load(); err != nil {
// 		logger.Error("No .env file found")
// 	}
// }

// func main() {
// 	err := logger.InitLogger("./logs", "app_logs.db")
// 	if err != nil {
// 		log.Fatalf("Failed to initialize logger: %v", err)
// 	}
// 	initConfig()

// 	partials := []string{
// 		"cmd/web/html/layouts/footer.html",
// 		"cmd/web/html/layouts/header.html",
// 		"cmd/web/html/layouts/navbar.html",
// 	}
// 	logger.Info("Starting server...")

// 	baseHandler := base.NewBaseHandler(
// 		"cmd/web/html/layouts/base.html",
// 		partials...,
// 	)

// 	queries, blogDB := dbschema.ConnectCreateDoShitWithDB()
// 	defer blogDB.Close()

// 	blogService := blog.NewBlogService(queries)
// 	homeService := home.NewHomeService(blogService)
// 	uploadService := addpost.NewPostSerivce(*queries)

// 	homeHandler := home.NewHomeHandler(homeService, baseHandler)
// 	blogHandler := blog.NewBlogHandler(blogService, baseHandler)
// 	uploadHandler := addpost.NewPostHandler(uploadService, baseHandler)
// 	handlers := routes.AppHandlers{
// 		HomeHandler: homeHandler,
// 		BlogHandler: blogHandler,
// 		PostHandler: uploadHandler,
// 	}
// 	logger.Info("Handlers registered")

// 	mux := routes.SetupRoutes(handlers)
// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		port = "4200" // Default port if not specified in the environment
// 	}

// 	server := &http.Server{
// 		Addr:    ":" + port,
// 		Handler: mux,
// 	}

// 	go func() {
// 		if err := server.ListenAndServe(); err != http.ErrServerClosed {
// 			logger.Fatal("Could not listen on %s: %v", port, err)
// 		}
// 	}()

// 	logger.Info("Server is running on port %s", port)

// 	// Graceful shutdown
// 	quit := make(chan os.Signal, 1)
// 	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
// 	<-quit
// 	logger.Info("Shutting down server...")

// 	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
// 	defer cancel()
// 	if err := server.Shutdown(ctx); err != nil {
// 		logger.Fatal("Server forced to shutdown: %v", err)
// 	}
// 	logger.Info("Server shutdown gracefully")
// }
