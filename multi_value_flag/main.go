package main

import (
	"flag"
	"fmt"
	"strings"
)

type MultiValueFlag []string

func (m *MultiValueFlag) String() string {
	return strings.Join(*m, ",")
}

func (m *MultiValueFlag) Set(value string) error {
	*m = append(*m, value)
	return nil
}

func main() {
	var f MultiValueFlag

	flag.Var(&f, "flag", "A multi-value flag")

	flag.Parse()

	fmt.Println(f)
}
