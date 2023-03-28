package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func main() {
	// Read the first YAML file
	path1 := os.Args[1]
	file1, err := ioutil.ReadFile(path1)
	if err != nil {
		log.Fatalf("Failed to read %s: %v", path1, err)
	}

	// Read the second YAML file
	path2 := os.Args[2]
	file2, err := ioutil.ReadFile(path2)
	if err != nil {
		log.Fatalf("Failed to read %s: %v", path2, err)
	}

	// Merge the two YAML files
	var merged map[interface{}]interface{}
	err = yaml.Unmarshal(append(file1, file2...), &merged)
	if err != nil {
		log.Fatalf("Failed to merge YAML files: %v", err)
	}

	// Print the merged YAML
	out, err := yaml.Marshal(merged)
	if err != nil {
		log.Fatalf("Failed to marshal merged YAML: %v", err)
	}
	fmt.Println(string(out))
}
