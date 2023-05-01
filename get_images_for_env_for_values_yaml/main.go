package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	dotEnvVars, err := godotenv.Read(".env")
	handleError(err)

	out := []string{}

	for name, value := range dotEnvVars {
		if strings.Contains(name, "IMAGE") {
			out = append(out, fmt.Sprintf("%s=%s", name, value))
		}
	}

	for _, line := range out {
		fmt.Println(line)
	}
}

func handleError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
