package main

import "fmt"

func main()	{
	var name string = "Alice"
	var age int = 30
	var isStudent bool = false
	fmt.Printf("Name: %s, Age: %d, Is Student: %t\n", name, age, isStudent)

	// Short variable declaration
	city := "New York"
	fmt.Printf("City: %s\n", city)

	// Multiple variable declaration
	var (
		country = "USA"
		zipCode = 10001
	)
	fmt.Printf("Country: %s, Zip Code: %d\n", country, zipCode)
}
