package main

import "fmt"

func main() {
	a := 7
	b := &a

	fmt.Println(a,b)
	fmt.Printf("%T\n",b)

	// Use * to get value
	fmt.Println(*b)

	// change value with pointer
	*b = 10
	fmt.Println(a)
}
