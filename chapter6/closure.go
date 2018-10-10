package main

import "fmt"

func main() {
	// Closure
	add := func(x, y int) int {
		return x + y
	}

	fmt.Println(add(1, 1))
}
