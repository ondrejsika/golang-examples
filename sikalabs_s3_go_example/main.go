package main

import (
	"fmt"
	"log"
	"time"

	sikalabs_s3_go "github.com/sikalabs-go/sikalabs-s3-go"
	"github.com/sikalabsx/sikalabs-encrypted-go/pkg/encrypted"
)

func main() {
	s3ConfigFromEncrypted, err := encrypted.GetConfigSikaLabsEncryptedBucket1()
	handleError(err)
	s3Config := sikalabs_s3_go.S3Config{
		AccessKey:  s3ConfigFromEncrypted.AccessKey,
		SecretKey:  s3ConfigFromEncrypted.SecretKey,
		Region:     s3ConfigFromEncrypted.Region,
		BucketName: s3ConfigFromEncrypted.BucketName,
	}

	fmt.Println("Bucket Name:", s3Config.BucketName)

	// Create file
	sikalabs_s3_go.PutObject(
		s3Config,
		"sikalabs_s3_go_test.txt",
		[]byte("sikalabs-go/sikalabs-s3-go "+time.Now().Format("2006-01-02 15:04:05")),
	)

	// Get file
	data, err := sikalabs_s3_go.GetObject(s3Config, "sikalabs_s3_go_test.txt")
	handleError(err)
	fmt.Println("File Content (sikalabs_s3_go_test.txt):", string(data))
}

func handleError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
