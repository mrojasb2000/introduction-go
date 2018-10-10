package main

import "fmt"

// Convert Fahrenheit to Celsius
func main() {
	fmt.Println("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := (input - 32) * 5 / 9

	fmt.Println(output)
}
