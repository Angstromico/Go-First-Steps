package main

import "fmt"

//middlename := "Bob" // This will cause an error because it's not inside a function
var middlename string = "Anton"	// This is a package-level variable, but it's generally better to declare variables inside functions for better scope management

func main()	{
	var middlename string = "Bob" // This is now inside a function
	var name string = "Alice"
	var age int = 30
	var isStudent bool = false
	fmt.Printf("Name: %s, Middle Name: %s, Age: %d, Is Student: %t\n", name, middlename, age, isStudent)

	// Short variable declaration
	city := "New York"
	fmt.Printf("City: %s\n", city)

	// Multiple variable declaration
	var (
		country = "USA"
		zipCode = 10001
	)
	fmt.Printf("Country: %s, Zip Code: %d\n", country, zipCode)

	//Default values
	var defaultString string
	var defaultInt int
	var defaultBool bool
	fmt.Printf("Default String: '%s', Default Int: %d, Default Bool: %t\n", defaultString, defaultInt, defaultBool)

	// Scoped variables
	if age > 18 {
		var isAdult bool = true //isAdult is only available within this if block for scope	
		fmt.Printf("Is Adult: %t\n", isAdult)
		calculatedArea := calculateArea(5.0)
		fmt.Printf("Area of circle with radius 5: %.2f\n", calculatedArea)
	}
}

func calculateArea(radius float64) float64 {
	const pi = 3.14159 //pi is only available within this function for scope
	return pi * radius * radius
}
