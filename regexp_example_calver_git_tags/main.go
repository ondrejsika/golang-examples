package main

import (
	"fmt"
	"regexp"
)

func main() {
	testData := []string{
		"v2022.1.1",
		"v2022.10.1",
		"v2022.1.111",
		"v2022.10.111",
		"v2022.10",
		"v2022.10.111-shit",
	}
	for _, s := range testData {
		r := regexp.MustCompile(`^v\d{4}.\d{1,2}.(\d+)$`)
		match := r.FindStringSubmatch(s)
		ok := len(match) == 2
		if ok {
			val := match[1]
			fmt.Println(s, ok, val)
		} else {
			fmt.Println(s, ok)
		}
	}
}
