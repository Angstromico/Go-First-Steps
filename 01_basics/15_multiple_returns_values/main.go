package main

import "fmt"

func main() {
	// Example of a function that returns multiple values
	name, age := getPersonInfo()
	fmt.Printf("Name: %s, Age: %d\n", name, age)
}

func getPersonInfo() (string, int) {
	return "Alice", 30
}
