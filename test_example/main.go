package main

import (
	"fmt"

	"test_example/pkg/hello"
)

func main() {
	fmt.Println(hello.SayHello("Dela"))
	fmt.Println(hello.SayHello("Nela"))
}
