package main

import "fmt"

func main() {
	// var
	// var name = "Harsh"
	var age int32 = 19

	// const
	const isBool = true

	// Shorthand(inside a function)
	// name:= "Harsh"
	// email:= "jainharsh270@gmail.com"
	name,email := "Harsh","jainharsh270@gmail.com"
	size:= 1.3  // float64 default


	fmt.Println(name,age,isBool,size,email) 

	// get the type
	fmt.Printf("%T\n",size)

}
