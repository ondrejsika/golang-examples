package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
)

func deriveKey(password string) []byte {
	hash := sha256.Sum256([]byte(password))
	return hash[:]
}

func main() {
	password := "password"
	key := deriveKey(password)
	plaintext := []byte(`            ___________________
           < This is encrypted >
            -------------------
           /
 /)-_-(\  /
  (o o)
   \o/\__-----.
    \      __  \
     \| /_/  \ /\__/
      ||      \\
      ||      //
      /|     /|`)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err)
	}

	ciphertext := aesGCM.Seal(nil, nonce, plaintext, nil)
	full := append(nonce, ciphertext...)
	encoded := base64.StdEncoding.EncodeToString(full)

	fmt.Println("Encrypted (base64):", encoded)
}
