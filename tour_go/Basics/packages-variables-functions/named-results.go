package main 

import "fmt"
// Named results. Define variables as seen below
func split(sum int) (x, y int) {
	x = sum * 4/9
	y = sum -x 
	// Naked return, which is only used in short functions
	return
}

func main() {
	fmt.Println(split(12))
}