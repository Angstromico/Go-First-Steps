package main

import "fmt"

func main() {
	// Example of a variadic function
	fmt.Println("Sum:", sum(1, 2, 3, 4, 5))
}

// A variadic function can accept zero or more arguments of a specific type
// ... is an Ellipsis
func sum(numbers ...int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}
