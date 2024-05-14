package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func helloWorld() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("(helloWorld) GET %s\n", r.URL.Path)
		fmt.Fprintf(w, "Hello World on port 8000!")
	})
	HelloWorldServer.Handler = mux
	HelloWorldServer.Addr = ":8000"
	log.Println("Server starting on 0.0.0.0:8000, see http://127.0.0.1:8000")
	if err := HelloWorldServer.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalln(err)
	}
}

func ahojSvete() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("(ahojSvete) GET %s\n", r.URL.Path)
		fmt.Fprintf(w, "Ahoj Svete z portu 8001!")
	})
	AhojSveteServer.Addr = ":8001"
	AhojSveteServer.Handler = mux
	log.Println("Server starting on 0.0.0.0:8001, see http://127.0.0.1:8001")
	if err := AhojSveteServer.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalln(err)
	}
}

var HelloWorldServer http.Server
var AhojSveteServer http.Server

func main() {
	// Start servers in their own goroutines to handle requests
	go helloWorld()
	go ahojSvete()

	// Setting up signal capturing
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until a signal is received
	<-stop

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	log.Println("Shutting down gracefully, press Ctrl+C again to force")
	if err := HelloWorldServer.Shutdown(ctx); err != nil {
		log.Fatalf("Could not gracefully shutdown the server: %v\n", err)
	}
	if err := AhojSveteServer.Shutdown(ctx); err != nil {
		log.Fatalf("Could not gracefully shutdown the server: %v\n", err)
	}
}
