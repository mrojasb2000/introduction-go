package main

import "fmt"

func main() {
	const x string = "Hello, World"
	x = "Some other string" //Error: cannot assign to x
	fmt.Println(x)
}
