package main

import "fmt"

// An Array is a numered sequuence of elements of a single type with a fixed length.
func main() {
	var x [5]int
	x[4] = 100
	fmt.Println(x) // result: [0 0 0 0 100]
	fmt.Println(x[4])
}
