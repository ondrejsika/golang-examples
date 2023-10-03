package main

import (
	"bytes"
	"fmt"
	"image"
	"log"

	"github.com/ondrejsika/go-dela"
	"github.com/qeesung/image2ascii/convert"

	_ "image/jpeg"
)

func main() {
	img, _, err := image.Decode(bytes.NewReader(dela.DELA1_JPG))
	if err != nil {
		log.Fatalln(err)
	}
	converter := convert.NewImageConverter()
	fmt.Print(converter.Image2ASCIIString(img, &convert.DefaultOptions))
}
