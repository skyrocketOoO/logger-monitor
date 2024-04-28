package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type LogEntry struct {
	Timestamp time.Time `json:"timestamp"`
	Line      string    `json:"line"`
}

func main() {
	logEntry := LogEntry{
		Timestamp: time.Now(),
		Line:      "Hello, world!",
	}
	logEntryBytes, err := json.Marshal(logEntry)
	if err != nil {
		fmt.Println("Error marshaling log entry:", err)
		return
	}

	url := "http://localhost:3102/loki/api/v1/push"

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(logEntryBytes))
	if err != nil {
		fmt.Println("Error sending log entry:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Log entry sent successfully")
}
