package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

func main() {
	// Load AWS config (credentials, region, etc.)
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("failed to load AWS config: %v", err)
	}

	// Create STS client
	client := sts.NewFromConfig(cfg)

	// Call GetCallerIdentity
	resp, err := client.GetCallerIdentity(context.TODO(), &sts.GetCallerIdentityInput{})
	if err != nil {
		log.Fatalf("failed to get caller identity: %v", err)
	}

	// Output the results
	fmt.Println("Account: ", *resp.Account)
	fmt.Println("User ID: ", *resp.UserId)
	fmt.Println("ARN:     ", *resp.Arn)
}
