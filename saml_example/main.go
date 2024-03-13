package main

import (
	"context"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/crewjam/saml/samlsp"
)

func main() {
	// openssl req -x509 -newkey rsa:2048 -keyout key.pem -out crt.pem -days 3650 -nodes -subj "/CN=example"
	keyPair, err := tls.LoadX509KeyPair("crt.pem", "key.pem")
	if err != nil {
		log.Fatalf("error loading key pair: %v", err)
	}

	key := keyPair.PrivateKey.(*rsa.PrivateKey)
	cert, err := x509.ParseCertificate(keyPair.Certificate[0])
	if err != nil {
		log.Fatalf("error parsing certificate: %v", err)
	}

	rootURL := "http://localhost:8000"

	idpMetadataURL, err := url.Parse("https://sso.sikademo.com/realms/sikademo/protocol/saml/descriptor")
	if err != nil {
		log.Fatalf("error parsing url: %v", err)
	}
	idpMetadata, err := samlsp.FetchMetadata(context.Background(), http.DefaultClient,
		*idpMetadataURL)
	if err != nil {
		log.Fatalf("error fetching metadata: %v", err)
	}

	// Parse the root URL
	parsedRootURL, err := url.Parse(rootURL)
	if err != nil {
		log.Fatalf("error parsing root URL: %v", err)
	}

	// Create SAML SP middleware
	samlSP, err := samlsp.New(samlsp.Options{
		EntityID:    "saml-example",
		URL:         *parsedRootURL,
		Key:         key,
		Certificate: cert,
		IDPMetadata: idpMetadata,
		SignRequest: true,
	})
	if err != nil {
		log.Fatalf("error creating SAML SP: %v", err)
	}

	app := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username := samlsp.AttributeFromContext(r.Context(), "username")
		email := samlsp.AttributeFromContext(r.Context(), "email")
		fmt.Fprintf(w, "Hello %s %s!\n", username, email)
	})

	http.Handle("/", samlSP.RequireAccount(app))
	http.Handle("/saml/", samlSP)

	log.Println("Server started at", rootURL)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
