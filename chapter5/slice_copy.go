package main

import "fmt"

// If the lengths of the two slices are not the same,
// the smaller of the two will be used.
func main() {
	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 2) // Create new slice only first two elements
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)
}
