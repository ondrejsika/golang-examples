package main

import (
	"bytes"
	"log"
	"os"
	"os/exec"
)

func Exec(user, password, host, command string) error {
	c := exec.Command(
		"sshpass", "-p", password,
		"ssh", user+"@"+host,
		"sudo", "-S", command,
	)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Stdin = bytes.NewBufferString(password + "\n")
	return c.Run()
}

func main() {
	password := os.Args[1]
	user := os.Args[2]
	host := os.Args[3]

	err := Exec(user, password, host, "id")
	if err != nil {
		log.Fatal(err)
	}
}
