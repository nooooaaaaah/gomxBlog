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
	file, err := os.OpenFile("tmp/blog.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	LogInfo = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	LogError = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile|log.Lmsgprefix)
}
