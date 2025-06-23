package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	baseURL := flag.String("keycloak-url", "", "Keycloak base URL (e.g. https://sso.sikalabs.com)")
	adminUser := flag.String("admin-username", "", "Admin username")
	adminPass := flag.String("admin-password", "", "Admin password")
	realm := flag.String("realm", "", "Keycloak realm name")
	username := flag.String("username", "", "Username to reset password for")
	password := flag.String("password", "", "New password")

	flag.Parse()

	if *baseURL == "" || *realm == "" || *adminUser == "" || *adminPass == "" || *username == "" || *password == "" {
		fmt.Println("Missing required arguments. Use -h for help.")
		os.Exit(1)
	}

	token, err := getToken(*baseURL, *realm, *adminUser, *adminPass, "admin-cli")
	if err != nil {
		panic("Token error: " + err.Error())
	}

	userID, err := getUserID(*baseURL, *realm, token, *username)
	if err != nil {
		panic("User lookup error: " + err.Error())
	}

	err = resetPassword(*baseURL, *realm, token, userID, *password, true)
	if err != nil {
		panic("Password reset failed: " + err.Error())
	}

	fmt.Println("Password reset successfully!")
}

func getToken(baseURL, realm, user, pass, clientID string) (string, error) {
	data := []byte(fmt.Sprintf("username=%s&password=%s&grant_type=password&client_id=%s", user, pass, clientID))

	resp, err := http.Post(
		fmt.Sprintf("%s/realms/master/protocol/openid-connect/token", baseURL, realm),
		"application/x-www-form-urlencoded",
		bytes.NewBuffer(data),
	)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	token, ok := result["access_token"].(string)
	if !ok {
		return "", fmt.Errorf("failed to get access token: %v", result)
	}
	return token, nil
}

func getUserID(baseURL, realm, token, username string) (string, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/admin/realms/%s/users?username=%s", baseURL, realm, username), nil)
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var users []map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&users)

	if len(users) == 0 {
		return "", fmt.Errorf("user not found")
	}

	return users[0]["id"].(string), nil
}

func resetPassword(baseURL, realm, token, userID, password string, temporary bool) error {
	passwordPayload := map[string]interface{}{
		"type":      "password",
		"value":     password,
		"temporary": temporary,
	}
	payloadBytes, _ := json.Marshal(passwordPayload)

	req, _ := http.NewRequest("PUT",
		fmt.Sprintf("%s/admin/realms/%s/users/%s/reset-password", baseURL, realm, userID),
		bytes.NewBuffer(payloadBytes),
	)
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to reset password: %s", body)
	}
	return nil
}
