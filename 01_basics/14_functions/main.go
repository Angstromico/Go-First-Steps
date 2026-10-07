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

	// Using the applyOperation function
	result := applyOperation(assignFunc, 5, 6)
	fmt.Println("Result of applyOperation:", result)

	// Using the getOperation function
	addFunc := getOperation("add")
	multiplyFunc := getOperation("multiply")

	fmt.Println("Result of addFunc:", addFunc(5, 6))
	fmt.Println("Result of multiplyFunc:", multiplyFunc(5, 6))
}

func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

func applyOperation(f func(int, int) int, x, y int) int {
	return f(x, y)
}

// Function that returns a function
func getOperation(operation string) func(int, int) int {
	switch operation {
	case "add":
		return func(x, y int) int {
			return x + y
		}
	case "multiply":
		return func(x, y int) int {
			return x * y
		}
	default:
		return nil
	}
}
