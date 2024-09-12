package server

import (
	"Blog/internal/routes"
	"Blog/pkg/logger"
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	port     string
	handlers routes.AppHandlers
}

func New(port string, handlers routes.AppHandlers) *Server {
	return &Server{
		port:     port,
		handlers: handlers,
	}
}

func (s *Server) Run() {
	mux := routes.SetupRoutes(s.handlers)

	server := &http.Server{
		Addr:    ":" + s.port,
		Handler: mux,
	}

	go func() {
		logger.Info("Server is running on port %s", s.port)
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			logger.Fatal("Could not listen on %s: %v", s.port, err)
		}
	}()

	s.gracefulShutdown(server)
}

func (s *Server) gracefulShutdown(server *http.Server) {
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
