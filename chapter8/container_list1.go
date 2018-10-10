package main

import (
	"container/list"
	"fmt"
)

func main() {
	var x list.List

	x.PushBack(1)
	x.PushBack(2)
	x.PushBack(3)

	fmt.Println("Front -> End")
	for e := x.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int))
	}
	fmt.Println("End -> Front")
	for e := x.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value.(int))
	}
}
