package main

import (
	"fmt"
	"regexp"
)

var MAKEFILE_WITH_UP_TARGET = `
IMAGE = example

up: build
	docker-compose up -d

build:
	docker-compose build

down:
	docker-compose down
`

var MAKEFILE_WITHOUT_UP_TARGET = `
nup:
	go build
`

func checkUpTargetInMakefile(makefile string) bool {
	r := regexp.MustCompile(`\nup:`)
	match := r.MatchString(makefile)
	return match
}

func main() {
	fmt.Println(
		"Makefile with up targer:",
		checkUpTargetInMakefile(MAKEFILE_WITH_UP_TARGET))
	fmt.Println(
		"Makefile without up targer:",
		checkUpTargetInMakefile(MAKEFILE_WITHOUT_UP_TARGET))
}
