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

func (e employee) presentation() {
	fmt.Printf("Hi, I'm %s and I'm %d years old, my employee ID is %s and my salary is %f \n", e.name, e.age, e.empId, e.salary)
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
