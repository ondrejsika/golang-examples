package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"

	"github.com/nfnt/resize"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <input-file> <output-file>")
		return
	}

	inputFileName := os.Args[1]
	outputFileName := os.Args[2]

	// Open the image file
	file, err := os.Open(inputFileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// Decode the image
	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Resize the image to fit within 256x256 pixels
	resizedImg := resize.Thumbnail(256, 256, img, resize.Lanczos3)

	// Create a new file to save the resized image
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer outputFile.Close()

	// Encode the resized image and save it to the file
	err = jpeg.Encode(outputFile, resizedImg, nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Image resized and saved successfully.")
}
