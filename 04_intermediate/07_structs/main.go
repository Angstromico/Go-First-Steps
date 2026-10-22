package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	firstName string
	lastName  string
	id        int
	age       int
}

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base
	str string
}

func main() {
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

	//Modify  a field within a struct
	p1.ModifyAge()

	fmt.Println(p1.age) // 5

	// Show full info from person:
	manuel := p1.ComputedPerson()
	fmt.Println(manuel)

	//Embedding structs inside other structs:
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)
	fmt.Println("also num:", co.base.num)
	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println("describer:", d.describe())
}

func (p *Person) ModifyAge() {
	p.age = 5
}

func (p *Person) ComputedPerson() string {
	return p.firstName + " " + strconv.Itoa(p.age)
}
