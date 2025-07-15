package main

import (
	"os"

	"github.com/sikalabs/hello-world-server/pkg/server"
)

func main() {
	os.Setenv("TEXT", "Hello Heureka!")
	os.Setenv("COLOR", "#0096FE")
	server.Server()
}
