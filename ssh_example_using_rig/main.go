package main

import (
	"fmt"

	"github.com/k0sproject/rig"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	rig.SetLogger(logger)

	conn := rig.Connection{
		SSH: &rig.SSH{
			User:    "root",
			Address: "lab0.sikademo.com",
			// PasswordCallback: func() (string, error) {
			// 	return "password", nil
			// },
		},
	}
	if err := conn.Connect(); err != nil {
		logger.Fatal(err)
	}
	defer conn.Disconnect()

	output, err := conn.ExecOutput("cat /etc/hostname")
	if err != nil {
		logger.Fatal(err)
	}
	fmt.Println("hostname", output)
}
