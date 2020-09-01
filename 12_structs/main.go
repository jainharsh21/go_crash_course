package main

import ("fmt"
		"strconv")

// Person struct definition
type Person struct{
	// firstName string
	// lastName string
	// city string
	// gender string
	// age int
	
	firstName,lastName,city,gender string
	age int
}

// Greeting method (Value reciever)
func (p Person) greet() string{
	return "Hello,my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (Pointer reciever)
func (p *Person) hasBirthday() {
	p.age++
} 

func main() {
	// init person using struct
	person1 := Person {firstName:"John",lastName:"Doe",city:"LA",gender:"M",age:25}
	// Alternative
	person2 := Person {"Mary","Jane","NYC","F",23}
	fmt.Println(person1)
	person2.age++
	fmt.Println(person2.age)

	person1.hasBirthday()
	fmt.Println(person1.greet())
}
