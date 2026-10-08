package main

import "fmt"

func main() {
	// Defer	is used to ensure that a function call is performed later in a program’s execution, usually for purposes of cleanup. defer is often used where e.g. ensure and finally would be used in other languages.
	defer fmt.Println("This will be printed last.")
	fmt.Println("This will be printed first.")
}
