package main

import (
	"fmt"
)

// Convert feet to meters.
func main() {
	fmt.Println("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 0.3084

	fmt.Println(output)
}
