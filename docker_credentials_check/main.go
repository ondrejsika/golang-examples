package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/remote/transport"
)

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Println("Usage: docker_credentials_check <registry>")
		return
	}
	registry := flag.Arg(0)

	err := checkRegistry(registry)
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
	} else {
		fmt.Println("✅ Success: Valid credentials")
	}
}

func checkRegistry(registry string) error {
	reg, err := name.NewRegistry(registry)
	if err != nil {
		return err
	}

	auth, err := authn.DefaultKeychain.Resolve(reg)
	if err != nil {
		return err
	}

	tr, err := transport.NewWithContext(
		context.Background(),
		reg,
		auth,
		http.DefaultTransport,
		[]string{reg.Scope(transport.PullScope)},
	)
	if err != nil {
		return err
	}

	client := &http.Client{Transport: tr, Timeout: 10 * time.Second}

	url := fmt.Sprintf("https://%s/v2/", reg.RegistryStr())
	resp, err := client.Do(&http.Request{Method: http.MethodGet, URL: mustParseURL(url)})
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		return nil
	}

	return fmt.Errorf("invalid or unauthorized: %s\n", resp.Status)
}

func mustParseURL(rawURL string) *url.URL {
	u, err := url.Parse(rawURL)
	if err != nil {
		panic(err)
	}
	return u
}
