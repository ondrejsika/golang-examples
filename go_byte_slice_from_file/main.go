package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: go run thisprogram.go filename")
	}

	filename := os.Args[1]
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	varName := strings.Split(filename, ".")[0]
	varName = strings.ReplaceAll(varName, "-", "_") // Replace hyphens with underscores

	fmt.Printf("var %s = []byte{\n", varName)
	for i, b := range data {
		fmt.Printf("0x%02x, ", b)
		if (i+1)%12 == 0 {
			fmt.Println()
		}
	}
	fmt.Print("\n}\n")
}
