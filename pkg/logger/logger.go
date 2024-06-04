package logger

import (
	"log"
	"os"
)

var (
	LogInfo  *log.Logger
	LogError *log.Logger
)

func init() {
	// Check if the file exists, if not, create it
	if _, err := os.Stat("tmp/blog.log"); os.IsNotExist(err) {
		file, err := os.Create("tmp/blog.log")
		if err != nil {
			log.Fatal(err)
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
