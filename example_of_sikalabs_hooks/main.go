package main

import (
	"fmt"
	"os"
	"os/exec"
)

func commandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

func execIfCommandExists(hook string) int {
	if commandExists(hook) {
		fmt.Println("info: run " + hook)
		cmd := exec.Command(hook)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println("error: " + err.Error())
			return 1
		} else {
			return 0
		}
	} else {
		fmt.Println("info: skip " + hook)
		return 0
	}
}

func main() {
	cmd := "_sikalabs_hook_example"
	execIfCommandExists(cmd)
}
