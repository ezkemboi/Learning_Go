package main

import "fmt"

func main() {
	var i int
	var s string
	var b bool
	var f float64
	// Use %q for string
	// String value will be ""
	// For boolean is False
	fmt.Printf("%v %q %v %v\n", i, s, b, f)
}