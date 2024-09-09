package logger

import (
	"log"
	"os"
	"path/filepath"
)

var (
	LogInfo  *log.Logger
	LogError *log.Logger
)

func init() {
	dirPath := "tmp"
	filePath := filepath.Join(dirPath, "blog.log")

	// Create the directory if it doesn't exist
	if err := os.MkdirAll(dirPath, 0755); err != nil {
		log.Fatalf("Could not create directory: %v", err)
	}

	// Check if the file exists, if not, create it
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			log.Fatalf("Could not create log file: %v", err)
		}
		file.Close()
	}

	// Open the file for writing
	file, err := os.OpenFile("tmp/blog.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize loggers
	LogInfo = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	LogError = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile|log.Lmsgprefix)
}
