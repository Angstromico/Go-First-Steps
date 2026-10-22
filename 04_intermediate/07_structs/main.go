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
	fmt.Println(p2.id)       //If you try to access a non defined field the return is 0

	//Anonymous Structs
	newCar := struct {
		make    string
		model   string
		mileage int
	}{
		make:    "Ford",
		model:   "Taurus",
		mileage: 200000,
	}

	fmt.Println(newCar.model)

	//Methods within a struct:
	type Calculator struct {
		Add      func(a, b int) int
		Subtract func(a, b int) int
	}

	c := Calculator{
		Add: func(a, b int) int {
			return a + b
		},
		Subtract: func(a, b int) int {
			return a - b
		},
	}

	fmt.Println(c.Add(10, 5))      // 15
	fmt.Println(c.Subtract(10, 5)) // 5
}
