package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	req, err := http.NewRequest("GET", os.Getenv("VAULT_ADDR")+"/v1/sys/storage/raft/snapshot", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("X-Vault-Token", os.Getenv("VAULT_TOKEN"))

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	out, err := os.Create("vault-backup.snap")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	io.Copy(out, resp.Body)
}
