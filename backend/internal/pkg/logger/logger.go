package logger

import (
    "log"
    "os"
)

// Init initializes the logger with custom settings
func Init() {
    log.SetOutput(os.Stdout)
    log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

// Info logs an info message
func Info(message string) {
    log.Printf("[INFO] %s", message)
}

// Error logs an error message
func Error(message string) {
    log.Printf("[ERROR] %s", message)
}

// Warn logs a warning message
func Warn(message string) {
    log.Printf("[WARN] %s", message)
}