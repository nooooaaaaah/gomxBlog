package main

import (
	"Blog/internal/base"
	"Blog/internal/home"
	"Blog/pkg/logger"
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)
func init() {
    if err := godotenv.Load(); err != nil {
        logger.LogError.Println("No .env file found")
    }
}
func main() {
	partials := []string{
		"ui/html/layouts/footer.html",
		"ui/html/layouts/header.html",
		"ui/html/partials/sidebar.html",
	}
	logger.LogInfo.Println("starting")
	baseHandler := base.NewBaseHandler(
		"ui/html/layouts/base.html",
		partials...,
	)
	logger.LogInfo.Println("created basehandler")
	handlers := AppHandlers{
		baseHandler,
		home.NewHomeHandler(baseHandler),
	}
	logger.LogInfo.Println("registered handlers")
	mux := setupRoutes(handlers)
	logger.LogInfo.Println("routes are setup")

	port := os.Getenv("PORT")
      if port == "" {
          port = "4200" // default port

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.LogError.Fatalf("Could not listen on %s: %v\n", port, err)
		}
	}()

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
}
}
