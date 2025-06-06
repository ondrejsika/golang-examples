package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
)

// Password to derive key
var password = "password"

// Encrypted data: base64-encoded (nonce + ciphertext)
var encryptedDataBase64 = "Af7diD5Tz4EfEHNkPeE/XHRQlGqzY2OAXMzyJ3Gx5MkRW8X4yaXna1KdNyChybqMFGzA0S2xuMbLK6GLFbP8s9mLoYV6wH8dLxFez7mUgIVw+vR1/Ea4ACY9xWOrUB3VWVPXQpe0YfRCp327Gp1yxMmYNWIXqS6Huun5HDAyGAl9I6dGloaJpmm/+mgORay7xX4BMWaAaFHpYq7O1AzVsEA/f+R4P7OaqeogKsEACLHt3EIhsDv4YVc++3CaZeytEDC++4gFvk1f7RIrB++nUA27w7UVRM6AjvuD0wZEwYyP6CQM6ckoQpCw45RVdNVwmjwSLKcTRplMtPAorCKydZ+c+64M"

func deriveKey(password string) []byte {
	hash := sha256.Sum256([]byte(password))
	return hash[:]
}

func main() {
	key := deriveKey(password)

	encryptedData, err := base64.StdEncoding.DecodeString(encryptedDataBase64)
	if err != nil {
		log.Fatalf("Failed to decode encrypted data: %v", err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalf("Failed to create cipher: %v", err)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatalf("Failed to create GCM: %v", err)
	}

	nonceSize := aesGCM.NonceSize()
	if len(encryptedData) < nonceSize {
		log.Fatal("Ciphertext too short")
	}

	nonce, ciphertext := encryptedData[:nonceSize], encryptedData[nonceSize:]

	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		log.Fatalf("Decryption failed: %v", err)
	}

	fmt.Println(string(plaintext))
}
