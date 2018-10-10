package main

import (
	"fmt"
)

// Print number between 1 to 100, bu replace multiples of 3 with
// "Fizz", multiples of 5 with "Buzz" and multiples of both 3 and 5 with "FizzBuzz"
func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
