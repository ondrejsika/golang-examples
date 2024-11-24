package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/robfig/cron/v3"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("Usage: %s <cron-expression> <command>", os.Args[0])
	}

	cronExpression := os.Args[1]
	command := os.Args[2]

	c := cron.New()
	_, err := c.AddFunc(cronExpression, func() {
		cmd := exec.Command("sh", "-c", command)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			log.Printf("Error executing command: %v", err)
		}
	})
	if err != nil {
		log.Fatalf("Failed to add cron job: %v", err)
	}

	log.Printf("Cron scheduler started with expression: %s, command: %s", cronExpression, command)
	c.Start()

	// Keep the program running
	select {}
}
