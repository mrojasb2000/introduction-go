package main

import (
	"flag"
	"fmt"
	"math/rand"
)

// go run command_args.go -max=9
// default max value is 6
func main() {

	// Define flags
	maxp := flag.Int("max", 6, "the max value")
	// Parse
	flag.Parse()
	// Generate a number between 0 on max
	fmt.Println(rand.Intn(*maxp))
}
