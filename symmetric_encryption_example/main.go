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
var password = "password"

// Encrypted data: base64-encoded (salt + nonce + ciphertext)
var encryptedDataBase64 = "DkhI5ZBYiASi/16kvpFCpL04Iv6ZkfV50nOixVTorrXDuc0d7wfz0DtCWq/vjhnP3bYLDfHCatUJ7JumOSM1n10IzYarJg1uAvv/lgkDElGNSQZzZ3qlUTphyYXReImjFm1DJCzg5Fbjin5Bp9l3NB0h024UUbYLAa5XUWG/QuDR3SC9KU91VBvIRQ5uCXPguUR6Biiu42KfPrkvZtJPBO/FAJQmZZ5Vd22pHTK++i50JpR1+dAQmCpEMwpUlYdt5okjGICu3qeeJqy6wlc1BhpVc8bzZ7wKUyPE3WfgWHcnVrfpvKq2ZouDa4AhCzeKNTVIpSbTxjmX5wzuCaSVk+bWdbM+1vgsWXKq9g1/pVufDNGMeg=="

func main() {
	encryptedData, err := base64.StdEncoding.DecodeString(encryptedDataBase64)
	if err != nil {
		log.Fatalf("Failed to decode encrypted data: %v", err)
	}

	// Extract salt (first 16 bytes)
	saltSize := 16
	if len(encryptedData) < saltSize {
		log.Fatal("Encrypted data too short to contain salt")
	}
	salt := encryptedData[:saltSize]
	remaining := encryptedData[saltSize:]

	// Derive key using salt
	key := pbkdf2.Key([]byte(password), salt, 100000, 32, sha256.New)

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalf("Failed to create cipher: %v", err)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatalf("Failed to create GCM: %v", err)
	}

	nonceSize := aesGCM.NonceSize()
	if len(remaining) < nonceSize {
		log.Fatal("Ciphertext too short")
	}

	nonce, ciphertext := remaining[:nonceSize], remaining[nonceSize:]

	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		log.Fatalf("Decryption failed: %v", err)
	}

	fmt.Println(string(plaintext))
}
