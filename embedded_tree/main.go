package main

import (
	"embed"
	"io"
	"io/fs"
	"log"
	"os"
	"path"
)

//go:embed data/*
//go:embed data/**/*
var srcDataFS embed.FS

func main() {
	var err error
	secFS, err := fs.Sub(srcDataFS, "data")
	handleError(err)
	fs.WalkDir(secFS, ".", func(filePath string, d fs.DirEntry, err error) error {
		handleError(err)
		if d.Type().IsDir() {
			os.MkdirAll(path.Join("out", filePath), 0755)
		} else {
			file, err := secFS.Open(filePath)
			handleError(err)
			fileContent, err := io.ReadAll(file)
			handleError(err)
			err = os.WriteFile(path.Join("out", filePath), fileContent, 0644)
			handleError(err)
		}
		return nil
	})
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
