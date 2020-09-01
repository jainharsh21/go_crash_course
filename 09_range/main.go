package main

import "fmt"

func main() {
	ids := []int{32,23,45,66,77,2}

	// loop through ids

	for i,id := range ids {
		fmt.Printf("%d - ID %d\n",i,id)
	}

	// without using index (_ is to not use the index)
	for _,id := range ids {
		fmt.Printf("ID : %d\n",id)
	}

	// Add IDs together
	sum := 0
	for _,id := range ids {
		sum += id
	}
	fmt.Println(sum)

	// range with map
	emails := map[string]string{"John":"john@gmail.com","Jane":"jane@gmail.com","Mike":"mike@gmail.com"}

	for k,v := range emails {
		fmt.Printf("%s : %s\n",k,v)
	}

	// get only keys
	for k := range emails {
		fmt.Printf("Name : %s\n",k)
	}

}
