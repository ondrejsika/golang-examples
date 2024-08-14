package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: go-cat <file1> [file2 ... fileN]")
		os.Exit(1)
	}

	for _, filename := range os.Args[1:] {
		err := catFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading file %s: %v\n", filename, err)
			os.Exit(1)
		}
	}
}

func catFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	_, err = io.Copy(os.Stdout, reader)
	if err != nil {
		return err
	}

	return nil
}

// created by chatgpt: https://chatgpt.com/share/3dc4b4e5-8e04-4123-aeaf-105cca5557d2
