package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

/*
func circleArea(c Circle) float64 {
	return math.Pi * c.r * c.r
}*/

/*
func circleArea(c *Circle) float64 {
	return math.Pi * c.r * c.r
}*/

// Method
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

// CaMethod
func (r *Rectangle) area() float64 {
	l := distance(x1, y1, x1, y2)
	w := distance(x1, y1, x2, y1)
	return l * w
}

func main() {
	// Variable local
	//var c Circle

	// Allocate memeory and return a pointer to the struct
	//c := new(Circle)

	// Inital value by fieldname
	//c := Circle{x: 0, y: 0, r: 5}

	// Initial value by order
	c := Circle{0, 0, 5}

	//fmt.Println(circleArea(c))

	// Callback by reference
	//fmt.Println(circleArea(&c))

	// Callback method with receiver
	fmt.Println(c.area())

	// Inital value return pointer
	//c := &Circle{0, 0, 5}

	//fmt.Println(c.x, c.y, c.r)
	//c.x = 10
	//c.y = 5
	//fmt.Println(c.x, c.y, c.r)

}
