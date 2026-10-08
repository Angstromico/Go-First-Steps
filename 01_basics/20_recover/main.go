package main

import "fmt"

func main() {
	//Using recover to handle panic situations gracefully. The recover function allows us to regain control of a panicking goroutine and continue execution without crashing the program.
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	panic("This is a panic situation!") // This will cause a panic and be recovered
}
