package main

import "fmt"

// Variables with initializers
var x, y int = 23, 12 

func main() {
	// get other variables initialized
	var c, python, javascript = true, false, "no!"
	fmt.Println(x, y, c, python, javascript)
}