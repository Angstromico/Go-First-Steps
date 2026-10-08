package main

import "fmt"

func main() {
	//Panic can recieve a balue of any type panic(interface{}) and stops the ordinary flow of control and begins panicking. When the function F calls panic, execution of F stops, any deferred functions in F are executed normally, and then F returns to its caller. To the caller, F then behaves like a call to panic.
	// fmt.Println("This will be printed first.")
	// panic("This is a panic!")
	// fmt.Println("This will not be printed.")

	// Recover is a built-in function that regains control of a panicking goroutine. Recover is only useful inside deferred functions. During normal execution, a call to recover will return nil and have no other effect. If the current goroutine is panicking, a call to recover will capture the value given to panic and resume normal execution.
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	//processPanic(10) // Normal execution
	processPanic(-5) // This will cause a panic and be recovered
}

func processPanic(input int) {
	// Function to process panic
	if input < 0 {
		panic("Negative value received!")
	}
	fmt.Printf("Processing input: %d\n", input)
}
