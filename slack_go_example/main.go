package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	token := os.Getenv("SLACK_TOKEN")
	channel := os.Getenv("SLACK_CHANNEL")

	api := slack.New(token)

	_, _, err := api.PostMessage(
		channel,
		slack.MsgOptionText("Hello from slack-go!", false),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error sending message: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("Message sent successfully")
}
