package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"

	"golang.org/x/crypto/pbkdf2"
)

// Password to derive key
const PASSWORD = "password"

// Encrypted data: base64-encoded (salt + nonce + ciphertext)
const ENCRYPTED = "TQcjWsRFpxJMpkmvPkQ5HjExITafJh1FJNpjvXwci/CBs9T9PbyEvotcafu1KoYjJzjMpyJPwVXRRodY/27I2h+lNBrLG4HjwmbnVuHwyav09xVbY931btwaGdkSVsWNi7oSQlGYVpQkNZabma2bhZj8gPaPVrngi0OdddNsyNb7igESnhK3pPkv610YfucsurJcBDYB1kb1jy+Y1+I7r4LkgZDrCZVxeXfqVOzFfm7Xq419DiriR3G4TDChtz9KO2e4wKt9szHS0rHvam4vRHo6ZAS/ex36XeTj2QKDGwsPGK1E8I+914z88SOhmUqBguaEXasrrrNb6xj37mhOIT1HmQbQvR4C1gfWzC9IUvl1/ffoXA=="

func main() {
	text, err := decrypt(PASSWORD, ENCRYPTED)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(text)
}

func decrypt(password, encryptedBase64 string) (string, error) {
	encryptedData, err := base64.StdEncoding.DecodeString(encryptedBase64)
	if err != nil {
		return "", err
	}

	// Extract salt (first 16 bytes)
	saltSize := 16
	if len(encryptedData) < saltSize {
		return "", fmt.Errorf("encrypted data too short")
	}
	salt := encryptedData[:saltSize]
	remaining := encryptedData[saltSize:]

	// Derive key using salt
	key := pbkdf2.Key([]byte(password), salt, 100000, 32, sha256.New)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := aesGCM.NonceSize()
	if len(remaining) < nonceSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	nonce, ciphertext := remaining[:nonceSize], remaining[nonceSize:]

	text, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(text), nil
}
