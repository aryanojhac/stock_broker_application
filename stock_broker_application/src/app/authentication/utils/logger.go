package utils

import (
	"log"
	"os"
)

// InfoLogger is used for informational messages
var InfoLogger *log.Logger

// ErrorLogger is used for error messages
var ErrorLogger *log.Logger

// InitLogger initializes the loggers
func InitLogger() {
	// Info logs will go to stdout
	InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	// Error logs will go to stderr
	ErrorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
