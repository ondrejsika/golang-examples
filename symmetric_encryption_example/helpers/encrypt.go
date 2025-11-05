package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"

	"golang.org/x/crypto/pbkdf2"
)

func main() {
	password := "password"
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

	// Generate a random salt (16 bytes is standard)
	salt := make([]byte, 16)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		panic(err)
	}

	key := pbkdf2.Key([]byte(password), salt, 100000, 32, sha256.New)

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

	// Structure: salt + nonce + ciphertext
	full := append(salt, nonce...)
	full = append(full, ciphertext...)
	encoded := base64.StdEncoding.EncodeToString(full)

	fmt.Println("Encrypted (base64):", encoded)
}
