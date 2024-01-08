package main

import (
	"fmt"
	"time"
)

func sendNotification(user string, message string) {
	// Simulate sending a notification
	fmt.Printf("Notification sent to %s: %s\n", user, message)
}

func main() {
	// Schedule multiple notifications using goroutines
	users := []string{"UserA", "UserB", "UserC"}
	message := "New message for you!"

	i := 0

	for _, user := range users {

		go sendNotification(user, message)
		fmt.Println(i)
		i = i + 1

	}

	// Wait for notifications to complete
	time.Sleep(2 * time.Second)
}

// In this example, GoLang uses goroutines to send notifications to multiple users concurrently,
// which can be especially beneficial for sending notifications to a large user base.
