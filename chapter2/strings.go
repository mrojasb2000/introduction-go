package main

import "fmt"

func main() {
	// Finds the length of a string
	fmt.Println(len("Hello, World"))
	fmt.Println(len(`Hello, World`)) // backticks

	// Accesses a particular character in the string
	fmt.Println("Hello, World"[1])

	// Concatenates two strings together
	fmt.Println("Hello," + " World")

}
