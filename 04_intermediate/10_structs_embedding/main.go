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

func (p person) presentation() {
	fmt.Printf("Hi, I am %s and I'm %d years old \n", p.name, p.age)
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
	fmt.Println("Employee name: ", emp.name) //You can access values embedded directly
	emp.presentation()
}
