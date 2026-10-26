package main

import "fmt"

type person struct {
	name string
	age  int
}

type employee struct {
	person //Embedded struct
	empId  string
	salary float32
}

func main() {
	emp := employee{
		person: person{
			name: "Manuel",
			age:  36,
		},
		empId:  "3d12345",
		salary: 740.00,
	}

	fmt.Printf("The employee info is: %+v \n", emp)
}
