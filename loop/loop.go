package main

import "fmt"

func main() {
	// For Loop
	for i := 0; i < 5; i++ {
		fmt.Println("for loop Iteration i-", i)
	}

	// While Loop (Go doesn't have while loops; you can mimic them with for loops)
	j := 0
	for j < 5 {
		fmt.Println("while loop Iteration j-", j)
		j++
	}

	// Range Loop (used for iterating over slices, maps, and channels)
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("Range Loop,Index: %d, Value: %d\n", index, value)
	}
}
