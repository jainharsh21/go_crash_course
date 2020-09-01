package main

import "fmt"

func main() {
	// // Arrays
	// var fruitArr [2]string
	// // Assign Values
	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Grapes"
	// Declare and assign
	// fruitArr := [2]string{"Apple","Grapes"}
	// fmt.Println(fruitArr)

	fruitSlice := []string{"Apple","Grapes","Watermelon"}
	fmt.Println(fruitSlice)
	fmt.Println(fruitSlice[1:2])
	fmt.Println(len(fruitSlice))

}
