package main


import "fmt"

/**
Simple report that calculates the current power capacity feeding grid
Compare it with current load and check on which it is utilized
*/

func main() {
	var plantcapacities []float64

	plantcapacities = []float64{30, 30, 30, 60, 60, 100}

	var capacity float64 = plantcapacities[0] + plantcapacities[1] + plantcapacities[2] + 
		plantcapacities[3] + plantcapacities[4] + plantcapacities[5]

	var gridLoad = 75.

	utilization := gridLoad / capacity

	fmt.Printf("%-20s%.0f\n", "Capacity: ", capacity)
	fmt.Printf("%-20s%.0f\n", "GridLoad: ", gridLoad)
	fmt.Printf("%-20s%.1f%%\n", "Utilization: ", utilization * 100)
}