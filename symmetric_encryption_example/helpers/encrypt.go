package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"log"

	"golang.org/x/crypto/pbkdf2"
)

const PASSWORD = "password"
const TEXT = `            ___________________
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
      /|     /|`

func main() {
	encrypted, err := encrypt(PASSWORD, TEXT)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(encrypted)
}

func encrypt(password, text string) (string, error) {
	// Generate a random salt (16 bytes is standard)
	salt := make([]byte, 16)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		return "", err
	}

	key := pbkdf2.Key([]byte(password), salt, 100000, 32, sha256.New)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := aesGCM.Seal(nil, nonce, []byte(text), nil)

	// Structure: salt + nonce + ciphertext
	full := append(salt, nonce...)
	full = append(full, ciphertext...)
	encoded := base64.StdEncoding.EncodeToString(full)

	return encoded, nil
}
