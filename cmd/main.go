package main

import (
	"Blog/internal/base"
	"Blog/internal/blog"
	"Blog/internal/home"
	"Blog/pkg/db"
	"Blog/pkg/logger"
	"Blog/routes"
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

func initConfig() {
	if err := godotenv.Load(); err != nil {
		logger.LogError.Println("No .env file found")
	}
}

func main() {
	initConfig()
	db.InitEdgeDB()
	defer db.CloseEdgeDB()
	partials := []string{
		"ui/html/layouts/footer.html",
		"ui/html/layouts/header.html",
		"ui/html/partials/sidebar.html",
	}
	logger.LogInfo.Println("Starting server...")

	baseHandler := base.NewBaseHandler(
		"ui/html/layouts/base.html",
		partials...,
	)
	blogService := blog.NewBlogService()
	homeService := home.NewHomeService(blogService)
	homeHandler := home.NewHomeHandler(homeService, baseHandler)
	blogHandler := blog.NewBlogHandler(blog.NewBlogService(), baseHandler)
	handlers := routes.AppHandlers{
		HomeHandler: homeHandler,
		BlogHandler: blogHandler,
	}
	logger.LogInfo.Println("Handlers registered")

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
			logger.LogError.Fatalf("Could not listen on %s: %v", port, err)
		}
	}()

	logger.LogInfo.Println("Server is running on port", port)

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.LogInfo.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.LogError.Fatalf("Server forced to shutdown: %v", err)
	}
	logger.LogInfo.Println("Server shutdown gracefully")
}
