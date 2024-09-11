package main

import (
	"Blog/dbschema"
	addpost "Blog/internal/addPost"
	"Blog/internal/base"
	"Blog/internal/blog"
	"Blog/internal/home"
	"Blog/pkg/logger"
	"Blog/routes"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

func initConfig() {
	if err := godotenv.Load(); err != nil {
		logger.Error("No .env file found")
	}
}

func main() {
	err := logger.InitLogger("./logs", "app_logs.db")
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	initConfig()

	partials := []string{
		"ui/html/layouts/footer.html",
		"ui/html/layouts/header.html",
		"ui/html/layouts/navbar.html",
	}
	logger.Info("Starting server...")

	baseHandler := base.NewBaseHandler(
		"ui/html/layouts/base.html",
		partials...,
	)

	queries, blogDB := dbschema.ConnectCreateDoShitWithDB()
	defer blogDB.Close()

	blogService := blog.NewBlogService(queries)
	homeService := home.NewHomeService(blogService)
	uploadService := addpost.NewPostSerivce(*queries)

	homeHandler := home.NewHomeHandler(homeService, baseHandler)
	blogHandler := blog.NewBlogHandler(blogService, baseHandler)
	uploadHandler := addpost.NewPostHandler(uploadService, baseHandler)
	handlers := routes.AppHandlers{
		HomeHandler: homeHandler,
		BlogHandler: blogHandler,
		PostHandler: uploadHandler,
	}
	logger.Info("Handlers registered")

	mux := routes.SetupRoutes(handlers)
	port := os.Getenv("PORT")
	if port == "" {
		port = "4200" // Default port if not specified in the environment
	}

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			logger.Fatal("Could not listen on %s: %v", port, err)
		}
	}()

	logger.Info("Server is running on port %s", port)

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.Fatal("Server forced to shutdown: %v", err)
	}
	logger.Info("Server shutdown gracefully")
}
