package main

import (
	_ "embed"
	"log"
	"os"
	"os/exec"
)

//go:embed values.yaml
var valuesYaml []byte

func main() {
	var err error
	tmpValuesFile, err := os.CreateTemp("", "helm-values-yaml-")
	handleError(err)
	defer os.Remove(tmpValuesFile.Name())
	_, err = tmpValuesFile.Write(valuesYaml)
	handleError(err)

	cmd := exec.Command(
		"helm",
		"upgrade", "--install",
		"hello-world",
		"hello-world", "--repo", "https://helm.sikalabs.io",
		"--values", tmpValuesFile.Name(),
	)
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		log.Fatalln(err)
	}
}

func handleError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
