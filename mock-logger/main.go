package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Start a goroutine to periodically generate log messages
	for {
		// Generate a mock log message
		logMessage := generateLogMessage()
		// Print the log message
		fmt.Println(logMessage)
		// Sleep for a random duration between 1 and 5 seconds
		sleepDuration := rand.Intn(5) + 1
		time.Sleep(time.Duration(sleepDuration) * time.Second)
	}

}

// Function to generate a mock log message
func generateLogMessage() string {
	// Current timestamp
	timestamp := time.Now().Format("2006-01-02T15:04:05.000Z07:00")
	// Mock log level
	logLevel := "INFO"
	// Mock log message
	message := "This is a mock log message"
	// Return the formatted log message
	return fmt.Sprintf("[%s] [%s] %s", timestamp, logLevel, message)
}
