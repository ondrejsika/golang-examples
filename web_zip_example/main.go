package main

import (
	"archive/zip"
	"io"
	"log"
	"net/http"
	"os"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	url := "https://releases.hashicorp.com/terraform/0.13.2/terraform_0.13.2_darwin_amd64.zip"
	WebZipToBin(url, "terraform", "terraform")
}

func WebZipToBin(url, inZipFileName, outFileName string) {
	var err error

	resp, err := http.Get(url)
	handleError(err)
	defer resp.Body.Close()

	tmpInFile, err := os.CreateTemp("", "go-zip-example")
	handleError(err)

	_, err = io.Copy(tmpInFile, resp.Body)
	handleError(err)

	r, err := zip.OpenReader(tmpInFile.Name())
	handleError(err)
	defer r.Close()

	for _, f := range r.File {
		if f.Name == inZipFileName {

			outFile, err := os.OpenFile(outFileName, os.O_CREATE|os.O_WRONLY, 755)
			handleError(err)
			defer outFile.Close()

			zipFile, _ := f.Open()
			defer zipFile.Close()

			_, err = io.Copy(outFile, zipFile)
			handleError(err)
		}
	}
}
