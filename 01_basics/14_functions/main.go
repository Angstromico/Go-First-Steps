package main

import "fmt"

func main() {
	// Example of a simple function
	greet("World")

	// Anonymous function example
	func() {
		fmt.Println("This is an anonymous function")
	}()

	//Assignment function example
	assignFunc := func(x, y int) int {
		return x + y
	}
	fmt.Println("Result of assignFunc:", assignFunc(3, 4))

	//Can also assign a no anonymous func to a variable:
	salutations := greet

	salutations("Alice")
}

func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}
