package main

import "fmt"

func main() {
	//var x map[string]int	//panic: runtime error: assignment to entry in nil map
	// Initialization before they can be used.
	/*
		x := make(map[string]int)
		x["key"] = 10
		fmt.Println(x["key"])
	*/

	x := make(map[int]int)
	x[1] = 10
	fmt.Println(x[1])
	fmt.Println("Map[int]int Length: ", len(x))
	delete(x, 1)
	fmt.Println("Map[int]int Length: ", len(x))
}
