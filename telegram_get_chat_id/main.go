package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")

	if botToken == "" {
		log.Fatalln("TELEGRAM_BOT_TOKEN must be set")
	}

	chatID, err := getLastChatID(botToken)
	if err != nil {
		log.Fatalf("Failed to get last chat ID: %v", err)
	}
	fmt.Println(chatID)
	sendMessage(botToken, int(chatID), fmt.Sprintf("`%d`", chatID))
}

func getLastChatID(botToken string) (int64, error) {
	type Chat struct {
		ID int64 `json:"id"`
	}

	type Message struct {
		Chat Chat `json:"chat"`
	}

	type Update struct {
		Message Message `json:"message"`
	}

	type Response struct {
		Ok     bool     `json:"ok"`
		Result []Update `json:"result"`
	}

	resp, err := http.Get("https://api.telegram.org/bot" + botToken + "/getUpdates")
	if err != nil {
		return 0, fmt.Errorf("getting updates failed: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("reading response body failed: %v", err)
	}

	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		return 0, fmt.Errorf("unmarshalling response failed: %v", err)
	}

	if len(response.Result) == 0 {
		return 0, fmt.Errorf("no updates found")
	}

	lastChatID := response.Result[len(response.Result)-1].Message.Chat.ID
	return lastChatID, nil
}

func sendMessage(botToken string, chatID int, message string) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)
	body, _ := json.Marshal(map[string]string{
		"chat_id":    fmt.Sprintf("%d", chatID),
		"text":       message,
		"parse_mode": "MarkdownV2",
	})
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("sending message failed: %v", err)
	}
	defer resp.Body.Close()
	return nil
}
