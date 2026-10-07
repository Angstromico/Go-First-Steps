package main

import "fmt"

func main() {
	// Example of a simple function
	greet("World")
}

func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}
