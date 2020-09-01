package main

import "fmt"

func main() {
	// // Define map
	// emails := make(map[string]string)
	
	// // Assign key-values
	// emails["John"] = "john@gmail.com"
	// emails["Jane"] = "jane@gmail.com"
	// emails["Mike"] = "mike@gmail.com"

	// Declare map and add key-values
	emails := map[string]string{"John":"john@gmail.com","Jane":"jane@gmail.com","Mike":"mike@gmail.com"}

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["John"])
	
	// Delete from map
	delete(emails,"Mike")
	fmt.Println(emails)

}
