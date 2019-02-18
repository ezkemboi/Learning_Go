package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println("X + Y = ", add(12, 12))
}
