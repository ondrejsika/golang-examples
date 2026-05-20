package main

import (
	"math/rand"
	"time"

	"github.com/sikalabs/hello-world-server/pkg/server"
)

func main() {
	colors := []string{"green", "yellow", "blue", "red", "black", "white"}
	go func() {
		for {
			server.Color = colors[rand.Intn(len(colors))]
			server.BackgroundColor = colors[rand.Intn(len(colors))]
			time.Sleep(1 * time.Second)
		}

	}()
	server.Color = "green"
	server.Server()
}
