package main

import "fmt"

func main() {
	//	Variable naming conventions
	//	1. Use camelCase for variable names
	//	2. Variable names should be descriptive and meaningful
	//	3. Avoid using underscores in variable names
	//	4. Use short variable names for loop counters and temporary variables

	var firstName string = "Alice"
	var lastName string = "Smith"
	var age int = 30
	var isStudent bool = false

	fmt.Printf("First Name: %s, Last Name: %s, Age: %d, Is Student: %t\n", firstName, lastName, age, isStudent)

	// PascalCase is typically used for exported variables (those that start with an uppercase letter)
	var FullName string = firstName + " " + lastName
	fmt.Printf("Full Name: %s\n", FullName)

	// Short variable names for loop counters
	for i := 0; i < 5; i++ {
		fmt.Printf("Loop iteration: %d\n", i)
	}
	//Range version
	for i := range 5 {
		fmt.Printf("Loop iteration: %d\n", i)
	}
	//Is also useful for temporary variables in functions
	area := calculateArea(5.0)
	fmt.Printf("Area of circle with radius 5: %.2f\n", area)

	//And structs, interfaces and enums are normally named using PascalCase as well
	type User struct {
		FirstName string
		LastName  string
		Age       int
		Active    bool
	}

	user := User{
		FirstName: "Alice",
		LastName:  "Smith",
		Age:       30,
		Active:    false,
	}

	fmt.Printf("User: %+v\n", user)
}

func calculateArea(radius float64) float64 {
	const pi = 3.14159
	return pi * radius * radius
}

// Constants examples
const maxRetries = 5
const DefaultTimeout = 30

//Package names should be lowercase and should not contain underscores or camelCase, keep them short and descriptive
//For example, a package for handling user authentication could be named "auth" or "authentication"
