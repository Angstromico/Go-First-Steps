package main

import "fmt"

func main() {
	type Person struct {
		firstName string
		lastName  string
		id        int
		age       int
	}

	//Initiallize struct:
	p := Person{
		firstName: "Manuel",
		lastName:  "Morales",
		id:        21018926,
		age:       36,
	}

	fmt.Printf("The person is: %v \n", p)
}
