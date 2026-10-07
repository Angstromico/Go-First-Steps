package main

import "fmt"

func main() {
	// Example of a simple function
	greet("World")

	// Anonymous function example
	func() {
		fmt.Println("This is an anonymous function")
	}()
}

func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}
