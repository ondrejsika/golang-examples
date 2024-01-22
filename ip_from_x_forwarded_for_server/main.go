package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ip := r.Header.Get("X-Forwarded-For")
		if ip == "" {
			ip = "X-Forwarded-For is empty"
		}
		fmt.Fprintf(w, "%s\n", ip)
	})

	fmt.Println("Listen on port http://0.0.0.0:8000")
	http.ListenAndServe(":8000", nil)
}
