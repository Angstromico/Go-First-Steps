package main

import "fmt"

func main() {
	//Pointers are variables that store the memory address of another variable.
	//In this example, we define a variable called "x" and assign it the value 10.
	//We then define a pointer variable called "p" and assign it the memory address of "x" using the "&" operator.
	//We can access the value of "x" using the pointer variable "p" by dereferencing it using the "*" operator.
	//Finally, we print the value of "x" and the value of "p" to demonstrate that they are both equal to 10.
	x := 10
	p := &x
	fmt.Println("Value of x:", x)  // Output: Value of x: 10
	fmt.Println("Value of p:", *p) // Output: Value of p: 10
}
