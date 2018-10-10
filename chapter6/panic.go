package main

import "fmt"

func main() {
	/*
		panic("PANIC")
		str := recover() // this will never happen
		fmt.Println(str)
	*/

	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")
}
