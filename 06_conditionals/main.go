package main

import "fmt"

func main() {
	x, y := 10, 10
	color := "red"
	// If,Else If,Else
	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else if x == y {
		fmt.Printf("%d is equal to %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	// Switch
	switch color {
	case "red" : 
		fmt.Println("Color is Red")
	case "blue" : 
		fmt.Println("Color is Blue")
	default : 
		fmt.Println("Color is neither Red nor Blue")
	}
}
