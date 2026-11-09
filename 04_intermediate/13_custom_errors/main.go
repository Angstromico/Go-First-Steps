package main

import "fmt"

type CustomError struct {
	Message string
}

func (e *CustomError) Error() string {
	return e.Message
}

func main() {
	err := &CustomError{Message: "This is a custom error"}
	fmt.Println(err)
}
