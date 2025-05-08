package main

import (
	"bytes"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	chatID := os.Getenv("TELEGRAM_CHAT_ID")
	filePath := os.Getenv("FILE_PATH")

	if botToken == "" || chatID == "" || filePath == "" {
		log.Fatal("Please set TELEGRAM_BOT_TOKEN, TELEGRAM_CHAT_ID, and FILE_PATH environment variables.")
	}

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// Add chat_id field
	_ = writer.WriteField("chat_id", chatID)

	// Add caption/message
	_ = writer.WriteField("caption", "Here is your file from Go!")

	// Add the file
	part, err := writer.CreateFormFile("document", filepath.Base(filePath))
	if err != nil {
		log.Fatalf("Failed to create form file: %v", err)
	}

	_, err = bytes.NewBuffer(nil).ReadFrom(file)
	if _, err = file.Seek(0, 0); err != nil {
		log.Fatalf("Failed to reset file: %v", err)
	}

	_, err = file.WriteTo(part)
	if err != nil {
		log.Fatalf("Failed to write file to part: %v", err)
	}

	writer.Close()

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendDocument", botToken)

	req, err := http.NewRequest("POST", url, &requestBody)
	if err != nil {
		log.Fatalf("Failed to create HTTP request: %v", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	fmt.Println("File sent. Status:", resp.Status)
}
