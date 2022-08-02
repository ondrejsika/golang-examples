package main

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
)

//go:embed static/*
var staticSource embed.FS

func main() {
	static, _ := fs.Sub(fs.FS(staticSource), "static")
	fs := http.FileServer(http.FS(static))
	http.Handle("/", fs)
	fmt.Println("See: http://127.0.0.1:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
