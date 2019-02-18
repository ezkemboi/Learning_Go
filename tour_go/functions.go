package main

import "fmt"

// First function to add two integers
func add(x int, y int) int {
	return x + y
}

func main() {
	// Use main to run add function.
	// Expected results is 24
	fmt.Println(add(12, 12))
}