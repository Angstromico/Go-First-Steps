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
	p1 := Person{
		firstName: "Manuel",
		lastName:  "Morales",
		id:        21018926,
		age:       36,
	}

	fmt.Printf("The person is: %v \n", p1)

	// You don't need to define allthe fields on an struct
	p2 := Person{
		firstName: "Yeri",
		lastName:  "Romero",
	}

	fmt.Println(p2.lastName) //Acces to a single field
}
