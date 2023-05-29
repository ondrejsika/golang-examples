package main

import (
	"fmt"
	"time"
)

func task(id int, sleep int, done chan<- int) {
	fmt.Printf("Task %d started\n", id)
	time.Sleep(time.Duration(sleep) * time.Second)
	fmt.Printf("Task %d finished\n", id)
	done <- id
}

func main() {
	const numJobs = 3

	done := make(chan int, numJobs)

	go task(1, 1, done)
	go task(2, 5, done)
	go task(3, 2, done)

	for i := 1; i <= numJobs; i++ {
		id := <-done
		fmt.Printf("Task %d is done\n", id)
	}

	fmt.Println("All tasks are done")
}
