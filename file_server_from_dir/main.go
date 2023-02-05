package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	fmt.Println("See: http://127.0.0.1:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
