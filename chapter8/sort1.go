package main

import (
	"fmt"
	"sort"
)

/*
  The sort function in sort package takes a sort.Interface
  and sorts it. The sort.Interface requires methods: Len, Less and Swap
*/
type Person struct {
	Name string
	Age  int
}

// By Name
type ByName []Person

func (ps ByName) Len() int {
	return len(ps)
}

func (ps ByName) Less(i, j int) bool {
	return ps[i].Name < ps[j].Name
}

func (ps ByName) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

// By Age
type ByAge []Person

func (ps ByAge) Len() int {
	return len(ps)
}

func (ps ByAge) Less(i, j int) bool {
	return ps[i].Age < ps[j].Age
}

func (ps ByAge) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func main() {
	kids := []Person{
		{"Jill", 12},
		{"Jack", 14},
	}
	// Sort by name
	sort.Sort(ByName(kids))
	fmt.Println(kids)

	// Sort by age
	sort.Sort(ByAge(kids))
	fmt.Println(kids)
}
