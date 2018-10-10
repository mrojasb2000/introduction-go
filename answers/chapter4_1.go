package main

import (
	"fmt"
)

// Print number between 1 and 100 that are divisible by 3
func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}
