package main

import "fmt"

func main() {
	for i := 0; i < 11; i++ {
		/*
			if i == 0 {
				fmt.Println("Zero")
			} else if i == 1 {
				fmt.Println("One")
			} else if i == 2 {
				fmt.Println("Two")
			} else if i == 3 {
				fmt.Println("Three")
			} else if i == 4 {
				fmt.Println("Four")
			} else if i == 5 {
				fmt.Println("Five")
			} else if i == 6 {
				fmt.Println("Six")
			} else if i == 7 {
				fmt.Println("Seven")
			} else if i == 8 {
				fmt.Println("Eigth")
			} else if i == 9 {
				fmt.Println("Nine")
			}*/

		switch i {
		case 0:
			fmt.Println("Zero")
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		case 2:
			fmt.Println("Three")
		case 4:
			fmt.Println("Four")
		case 5:
			fmt.Println("Five")
		case 6:
			fmt.Println("Six")
		case 7:
			fmt.Println("Seven")
		case 8:
			fmt.Println("Eight")
		case 9:
			fmt.Println("Nive")
		default:
			fmt.Println("Unknown Number")
		}
	}
}
