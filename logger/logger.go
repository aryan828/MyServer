package logger

import (
	"log"
	"os"
)

var logFile string = "/home/ubuntu/server.log"

// InitializeLogging to save server logs.
func InitializeLogging() *log.Logger {
	f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	// defer f.Close()
	return log.New(f, "->  ", log.LstdFlags)
}
