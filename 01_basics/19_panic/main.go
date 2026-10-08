package main

import "fmt"

func main() {
	//Panic can recieve a balue of any type panic(interface{}) and stops the ordinary flow of control and begins panicking. When the function F calls panic, execution of F stops, any deferred functions in F are executed normally, and then F returns to its caller. To the caller, F then behaves like a call to panic.
	fmt.Println("This will be printed first.")
	panic("This is a panic!")
	fmt.Println("This will not be printed.")
}
