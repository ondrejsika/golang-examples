package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type LicenseData struct {
	Name  string    `json:"name"`
	Email string    `json:"email"`
	From  time.Time `json:"from"`
	To    time.Time `json:"to"`
}

type License struct {
	LicenseData
	Signature string `json:"signature"`
}

func sign(privateKey ed25519.PrivateKey, data LicenseData) (string, error) {
	payload, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	sig := ed25519.Sign(privateKey, payload)
	return base64.StdEncoding.EncodeToString(sig), nil
}

func verify(publicKey ed25519.PublicKey, license License) error {
	sig, err := base64.StdEncoding.DecodeString(license.Signature)
	if err != nil {
		return fmt.Errorf("decode signature: %w", err)
	}
	payload, err := json.Marshal(license.LicenseData)
	if err != nil {
		return err
	}
	if !ed25519.Verify(publicKey, payload, sig) {
		return fmt.Errorf("invalid signature")
	}
	if time.Now().After(license.To) {
		return fmt.Errorf("license expired on %s", license.To.Format(time.DateOnly))
	}
	if time.Now().Before(license.From) {
		return fmt.Errorf("license not yet valid, starts %s", license.From.Format(time.DateOnly))
	}
	return nil
}

func main() {
	// Generate Ed25519 key pair
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Private key: %s\n", base64.StdEncoding.EncodeToString(privateKey))
	fmt.Printf("Public key:  %s\n\n", base64.StdEncoding.EncodeToString(publicKey))

	// Create and sign a license
	data := LicenseData{
		Name:  "John Doe",
		Email: "john@example.com",
		From:  time.Now().UTC().Truncate(24 * time.Hour),
		To:    time.Now().UTC().Truncate(24*time.Hour).AddDate(1, 0, 0),
	}

	sig, err := sign(privateKey, data)
	if err != nil {
		log.Fatal(err)
	}

	license := License{
		LicenseData: data,
		Signature:   sig,
	}

	licenseJSON, err := json.MarshalIndent(license, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("License JSON:\n%s\n\n", licenseJSON)

	// Verify the license
	if err := verify(publicKey, license); err != nil {
		fmt.Printf("Verification FAILED: %v\n", err)
	} else {
		fmt.Println("Verification OK: license is valid")
	}

	// Tamper with the license and verify again
	tampered := license
	tampered.Email = "hacker@evil.com"
	if err := verify(publicKey, tampered); err != nil {
		fmt.Printf("Tampered license verification FAILED (expected): %v\n", err)
	}
}
