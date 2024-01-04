package main

import (
	"fmt"
	"time"
)

func worker(id int) {
	for i := 0; i < 3; i++ {
		fmt.Printf("Worker %d: Job %d\n", id, i)
		time.Sleep(time.Second) // Simulate work with a pause
	}
}

func main() {
	for i := 0; i < 3; i++ {
		go worker(i) // Start three workers concurrently
	}

	// Wait a moment to let the workers finish
	time.Sleep(4 * time.Second)
}

//In this example, we have three workers (goroutines) doing their jobs concurrently, which makes our program more efficient and faster.
//Concurrency and goroutines are a unique feature of Go that simplifies the handling of multiple tasks and can greatly improve the performance of software.
