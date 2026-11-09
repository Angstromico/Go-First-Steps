package main

import "fmt"

type CustomError struct {
	Message string
	Code    int
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Error: %s (Code: %d)", e.Message, e.Code)
}

func main() {
	err := &CustomError{Message: "This is a custom error", Code: 1}
	fmt.Println(err)
}
