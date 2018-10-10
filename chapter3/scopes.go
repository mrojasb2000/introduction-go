package main

import "fmt"

// This means that other functions can access this variable.
var x string = "Hello, World"

func main() {
	//var x string = "Hello, World"
	fmt.Println(x)
	f()
}

func f() {
	fmt.Println("Greeting, ", x)
}
