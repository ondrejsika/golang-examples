package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	username := "username"
	password := "password"

	htpasswdEntry, err := generateHtpasswd(username, password)
	if err != nil {
		log.Fatalln("Error generating htpasswd entry:", err)
	}
	fmt.Println(htpasswdEntry)
}

func generateHtpasswd(username, password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s:%s", username, string(hashedPassword)), nil
}
