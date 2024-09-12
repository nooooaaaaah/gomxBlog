package config

import (
	"Blog/pkg/logger"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	// Add other configuration fields as needed
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		logger.Error("No .env file found")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "4200" // Default port if not specified in the environment
	}

	return &Config{
		Port: port,
		// Initialize other fields here
	}, nil
}
