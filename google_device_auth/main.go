package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"golang.org/x/oauth2"
)

func main() {
	CLIENT_ID := os.Getenv("CLIENT_ID")
	CLIENT_SECRET := os.Getenv("CLIENT_SECRET")
	SCOPES := []string{"https://www.googleapis.com/auth/drive.file"}

	ctx := context.Background()
	conf := &oauth2.Config{
		ClientID:     CLIENT_ID,
		ClientSecret: CLIENT_SECRET,
		Scopes:       SCOPES,
		Endpoint: oauth2.Endpoint{
			TokenURL:      "https://oauth2.googleapis.com/token",
			DeviceAuthURL: "https://oauth2.googleapis.com/device/code",
		},
	}

	code, err := conf.DeviceAuth(ctx)
	if err != nil {
		log.Fatalf("Failed to get device and user codes: %v", err)
	}

	fmt.Printf("Visit the URL: %s\n", code.VerificationURI)
	fmt.Printf("And enter the code: %s\n", code.UserCode)

	for {
		token, err := conf.DeviceAccessToken(ctx, code)
		if err == nil {
			fmt.Printf("Got access token: %s\n", token.AccessToken)
			fmt.Printf("Got refresh token: %s\n", token.RefreshToken)
			break
		}
		time.Sleep(1 * time.Second)
	}
}
