package main

import "fmt"

func main() {
	// Defer	is used to ensure that a function call is performed later in a program’s execution, usually for purposes of cleanup. defer is often used where e.g. ensure and finally would be used in other languages.
	defer fmt.Println("This will be printed last.")
	fmt.Println("This will be printed first.")

	multipleDefers(10)
}

// Function with multiple	defer statements
func multipleDefers(i int) {
	defer fmt.Println("The value is not modified after deferred statement: ", i)
	defer fmt.Println("First defer statement.")
	defer fmt.Println("Second defer statement.")
	defer fmt.Println("Third defer statement.")
	i++
	fmt.Println("This will be printed in the middle.")
}
