package main

import (
	"errors"
	"fmt"
)

func main() {
	// Example of handling an error
	_, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("denominator cannot be zero")
	}
	return a / b, nil
}
