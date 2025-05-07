package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"log"
)

func main() {
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	chatID := os.Getenv("TELEGRAM_CHAT_ID")

	if botToken == "" || chatID == "" {
		log.Fatalln("TELEGRAM_BOT_TOKEN and TELEGRAM_CHAT_ID must be set")
	}

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)

	message := map[string]string{
		"chat_id": chatID,
		"text":    "Hello from Go!",
	}

	body, _ := json.Marshal(message)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Message sent. Status:", resp.Status)
}
