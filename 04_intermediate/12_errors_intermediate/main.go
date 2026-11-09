package main

import "fmt"

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}

	return a / b, nil
}

// Function that only returns error:
func checkDivisible(a, b int) error {
	if a%b != 0 {
		return fmt.Errorf("%d is not divisible by %d", a, b)
	}
	return nil
}

func main() {
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	err = checkDivisible(10, 3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("10 is divisible by 3")
	}

	err = checkDivisible(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("10 is divisible by 2")
	}

	// Example usage of CustomError
	customErr := &CustomError{Message: "custom error message"}
	fmt.Println("Custom error:", customErr)
}

type CustomError struct {
	Message string
}

func (e *CustomError) Error() string {
	return e.Message
}

func (e *CustomError) Unwrap() error {
	return nil
}
