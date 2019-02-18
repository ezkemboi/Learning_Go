package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	// Swap to start with the second and then the first
	a, b := swap("Programming", "Go")
	fmt.Println("Result of X & Y swap:  ", a, b)

	// To print the first one first
	b, a = a, b
	fmt.Println("Result of X & Y swap:  ", a, b)
}
