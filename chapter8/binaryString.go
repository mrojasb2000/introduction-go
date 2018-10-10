package main

import (
	"fmt"
)

func main() {
	arr := []byte("test")
	//str := string([]byte{'t', 'e', 's', 't'})
	str := string(arr)
	fmt.Println(arr)
	fmt.Println(str)
}
