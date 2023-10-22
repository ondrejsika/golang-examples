package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"golang.org/x/oauth2"
	drive "google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

func main() {
	CLIENT_ID := os.Getenv("CLIENT_ID")
	CLIENT_SECRET := os.Getenv("CLIENT_SECRET")
	ACCESS_TOKEN := os.Getenv("ACCESS_TOKEN")

	FILE_TO_UPLOAD := "google_drive_upload_example.txt"
	SCOPES := []string{drive.DriveFileScope}

	conf := &oauth2.Config{
		ClientID:     CLIENT_ID,
		ClientSecret: CLIENT_SECRET,
		Scopes:       SCOPES,
		Endpoint: oauth2.Endpoint{
			TokenURL:      "https://oauth2.googleapis.com/token",
			DeviceAuthURL: "https://oauth2.googleapis.com/device/code",
		},
	}

	ctx := context.Background()

	token := &oauth2.Token{
		AccessToken: ACCESS_TOKEN,
	}
	client := conf.Client(ctx, token)

	srv, err := drive.NewService(ctx, option.WithHTTPClient(client))

	if err != nil {
		log.Fatalf("Unable to retrieve Drive client: %v", err)
	}

	f, err := os.Open(FILE_TO_UPLOAD)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer f.Close()

	fileMetadata := &drive.File{
		Name: filepath.Base(FILE_TO_UPLOAD),
	}
	file, err := srv.Files.Create(fileMetadata).Media(f).Do()
	if err != nil {
		log.Fatalf("Failed to create file: %s", err)
	}

	fmt.Printf("File ID: %s\n", file.Id)
}
