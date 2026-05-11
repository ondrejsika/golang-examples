package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"


	"github.com/redis/go-redis/v9"
)

func handleError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func getAzureToken() (string, error) {
	req, err := http.NewRequest("GET", "http://169.254.169.254/metadata/identity/oauth2/token", nil)
	if err != nil {
		return "", err
	}
	q := req.URL.Query()
	q.Add("api-version", "2018-02-01")
	q.Add("resource", "https://redis.azure.com")
	req.URL.RawQuery = q.Encode()
	req.Header.Set("Metadata", "true")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var result struct {
		AccessToken string `json:"access_token"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}
	if result.AccessToken == "" {
		return "", fmt.Errorf("empty access token in response: %s", body)
	}
	return result.AccessToken, nil
}

func oidFromToken(token string) (string, error) {
	parts := strings.Split(token, ".")
	if len(parts) < 2 {
		return "", fmt.Errorf("invalid JWT format")
	}
	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return "", fmt.Errorf("failed to decode JWT payload: %w", err)
	}
	var claims struct {
		OID string `json:"oid"`
	}
	if err := json.Unmarshal(payload, &claims); err != nil {
		return "", fmt.Errorf("failed to parse JWT payload: %w", err)
	}
	if claims.OID == "" {
		return "", fmt.Errorf("oid claim not found in token")
	}
	return claims.OID, nil
}

func main() {
	host := os.Getenv("REDIS")
	if host == "" {
		log.Fatalln("REDIS env var is required")
	}

	token, err := getAzureToken()
	handleError(err)

	oid, err := oidFromToken(token)
	handleError(err)

	rdb := redis.NewClient(&redis.Options{
		Addr:     host + ":6379",
		Username: oid,
		Password: token,
	})

	ctx := context.Background()

	val, err := rdb.Incr(ctx, "counter").Result()
	handleError(err)

	fmt.Println(val)
}
