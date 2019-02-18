package main

import (
	"fmt"
	"math"
)

func main() {
	// Exported names start with capital letter
	fmt.Println("The value of PI is---->  : ", math.Pi)
	/**
	This will give an error if given (Notice small letter in pi)
	fmt.Println("This is PI", math.pi)
	*/
}