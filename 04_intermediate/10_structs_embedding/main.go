package main

import "fmt"

type person struct {
	name string
	age  int
}

type employee struct {
	employeeInfo person //Embedded struct named field
	person              //Anonymous struct
	empId        string
	salary       float32
}

func (p person) presentation() {
	fmt.Printf("Hi, I am %s and I'm %d years old \n", p.name, p.age)
}

func (e employee) presentation() {
	fmt.Printf("Hi, I'm %s and I'm %d years old, my employee ID is %s and my salary is %f \n", e.employeeInfo.name, e.employeeInfo.age, e.empId, e.salary)
}

func main() {
	emp := employee{
		employeeInfo: person{
			name: "Manuel",
			age:  36,
		},
		empId:  "3d12345",
		salary: 740.00,
	}

	fmt.Printf("The employee info is: %+v \n", emp)
	fmt.Println("Employee name: ", emp.employeeInfo.name) //You can access values embedded directly
	emp.presentation()
}
