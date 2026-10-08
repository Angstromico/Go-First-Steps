package main

import "fmt"

func main() {
	// Example of a variadic function
	fmt.Println("Sum:", sum(1, 2, 3, 4, 5))

	greet("Hello", "Alice", "Bob", "Charlie")

	// Passing a slice to a variadic function
	numbers := []int{10, 20, 30, 40, 50}
	fmt.Println("Sum of slice:", sumSlice(numbers))

	// Using a slice as the second argument in a function with two variadic parameters
	names := []string{"Alice", "Bob", "Charlie"}
	ages := []int{25, 30, 35}
	showNamesAndAges(names, ages...)
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

// A variadic function can accept a normal argument and another variadic:
func greet(greeting string, names ...string) {
	for _, name := range names {
		fmt.Printf("%s, %s!\n", greeting, name)
	}
}

// The variadic argument must be the last parameter in the function signature. You can also pass a slice to a variadic function by using the ... operator. For example:
func sumSlice(numbers []int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}

// In order to use two variadic arguments	in a function, you can use a slice as the second argument. For example:
func showNamesAndAges(names []string, ages ...int) {
	for i, name := range names {
		if i < len(ages) {
			fmt.Printf("%s is %d years old.\n", name, ages[i])
		}
	}
}
