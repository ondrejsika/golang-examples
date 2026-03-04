package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	name := filepath.Base(os.Args[0])

	switch name {
	case "hello":
		fmt.Println("Hello World!")
	case "ahoj":
		fmt.Println("Ahoj Svete!")
	case "dela":
		fmt.Println("Haf, I'm Dela!")
	case "nela":
		fmt.Println("Woof, I'm Nela!")
	default:
		fmt.Fprintf(os.Stderr, "unknown binary name: %s, should be one of: hello, ahoj, dela, nela\n", name)
		os.Exit(1)
	}
}
