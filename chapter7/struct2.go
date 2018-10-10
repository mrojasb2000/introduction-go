package main

import (
	"fmt"
)

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

type Android struct {
	// Android has a Person
	Person // Person	// Anonymous fields
	Model  string
}

func main() {

	p := new(Person)
	p.Name = "Mavro"

	a := new(Android)

	// We access to Person struct on Android
	a.Person = *p
	//a.Person.Talk()

	// We can also call any Person methods directly
	a.Talk()
}
