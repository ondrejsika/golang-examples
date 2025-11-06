package main

import (
	"fmt"
	"log"
	"time"

	sikalabs_s3_go "github.com/sikalabs-go/sikalabs-s3-go"
	"github.com/sikalabsx/sikalabs-encrypted-go/pkg/encrypted"
)

func main() {
	s3Config, err := encrypted.GetConfigSikaLabsEncryptedBucket1()
	handleError(err)

	fmt.Println("Bucket Name:", s3Config.BucketName)

	sikalabs_s3_go.PutObject(
		sikalabs_s3_go.S3Config{
			AccessKey:  s3Config.AccessKey,
			SecretKey:  s3Config.SecretKey,
			Region:     s3Config.Region,
			BucketName: s3Config.BucketName,
		},
		"sikalabs_s3_go_test.txt",
		[]byte("sikalabs-go/sikalabs-s3-go "+time.Now().Format("2006-01-02 15:04:05")),
	)
}

func handleError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
