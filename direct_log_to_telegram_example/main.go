package main

import (
	"fmt"
	"log"

	"github.com/sikalabsx/direct-log-to-telegram/pkg/direct_log_to_telegram"
)

func main() {
	// Send a simple log message to Telegram
	err := direct_log_to_telegram.Log("Test direct-log-to-telegram from golang-examples!")
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}

	fmt.Println("Message sent successfully!")
}
