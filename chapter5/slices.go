package main

import "fmt"

// A Slice is a type built on top of an array

func main() {
	//var x []float64 			// Declare slice empty
	//x := make(float64,5)		// Declare slice of length 5
	//x := make(float64, 5 10)	// Declare slice of length 5 and the capacity 10
	arr := [5]float64{1, 2, 3, 4, 5}
	//x := arr[0:5]
	//y := arr[1:4]
	//y := arr[0:] 				// equals arr[0:5]
	//y := arr[:5]				// equals arr[0:5]
	//y := arr[:]				// equals arr[0:5]
	fmt.Println(y)

}
