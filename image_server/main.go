package main

import (
	_ "embed"
	"fmt"
	"net/http"
)

//go:embed dela.jpg
var imageOfDela []byte

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/jpeg")
		w.WriteHeader(http.StatusOK)
		w.Write(imageOfDela)
	})
	fmt.Println("Listen on 0.0.0.0:8000, see http://127.0.0.1:8000")
	http.ListenAndServe(":8000", nil)
}
