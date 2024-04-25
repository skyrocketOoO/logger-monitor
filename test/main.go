package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type LogEntry struct {
	Stream string            `json:"stream"`
	Values map[string]string `json:"values"`
}

func main() {
	logEntry := LogEntry{
		Stream: "myapp",
		Values: map[string]string{
			"level":   "info",
			"message": "Hello, world!",
		},
	}

	jsonData, err := json.Marshal(logEntry)
	if err != nil {
		log.Fatalf("Error marshalling JSON: %s", err)
	}

	resp, err := http.Post("http://localhost:3100/loki/api/v1/push", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Error sending log entry to Loki: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Unexpected status code: %d", resp.StatusCode)
	}

	log.Println("Log entry sent successfully")
}
