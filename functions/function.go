package main

import "fmt"

// Function to add two numbers
func add(a, b int) int {
	return a + b
}

func main() {
	result := add(5, 3)
	fmt.Println("5 + 3 =", result)
}
