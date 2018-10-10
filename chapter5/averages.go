package main

import "fmt"

func main() {
	var total float64 = 0

	/*
		var x [5]float64
		x[0] = 98
		x[1] = 93
		x[2] = 77
		x[3] = 82
		x[4] = 83

		for i := 0; i < len(x); i++ {
			total += x[i]
		}*/

	// shorter syntax
	x := [5]float64{98, 93, 77, 82, 83}

	for _, value := range x {
		total += value
	}

	fmt.Println("Average: ", total/float64(len(x)))
}
