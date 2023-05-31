package main

import (
	"flag"
	"fmt"
)

func main() {
	var hello string
	flag.StringVar(&hello, "hello", "Hello", "Hello in different language")

	name := flag.String("name", "World", "Name to say hello to")

	flag.Parse()

	fmt.Printf("%s %s!\n", hello, *name)
}
