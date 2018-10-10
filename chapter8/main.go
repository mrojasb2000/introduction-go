package main

import (
	"fmt"
	m "github.com/mrojasb2000/golang-book/chapter8/math"
)

// Before run application install package custom math.

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := m.Average(xs)
	fmt.Println(avg)
}
