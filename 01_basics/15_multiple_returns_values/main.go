package main

import "fmt"

func main() {
	// Example of a function that returns multiple values
	name, age := getPersonInfo()
	fmt.Printf("Name: %s, Age: %d\n", name, age)

	// Example of a function that returns multiple values including an error
	message, err := compare(5, 4)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Comparison result:", message)
	}
}

func getPersonInfo() (string, int) {
	return "Alice", 30
}

// A function that returns multiple variables can come in handy for debugging errors:
func compare(a, b int) (string, error) {
	if a == b {
		return "Values are equal", nil
	}
	return "Values are not equal", fmt.Errorf("values do not match")
}
