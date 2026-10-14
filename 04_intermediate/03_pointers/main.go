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
	fmt.Println("Direction on memory of the variable:", p)
	fmt.Println("Direction on memory of the variable:", &x)
	//Pointer of nil value:

	var p1 *int
	fmt.Println("Pointer of nil value:", p1) // Output: Pointer of nil value: <nil>

	//There is no pointer arithmetic in Go. You cannot perform operations like incrementing or decrementing a pointer value. However, you can use the "unsafe" package to perform pointer arithmetic, but it is not recommended as it can lead to undefined behavior.

	//In this example, we define a function called "modifyPointerValue" that takes a pointer to an integer as an argument.
	//Inside the function, we dereference the pointer using the "*" operator and assign it a new value of 20.
	//We then call the function and pass the pointer variable "p" as an argument.
	//Finally, we print the value of "x" to demonstrate that it has been modified to 20.
	modifyPointerValue(p)
	fmt.Println("Value of x after modification:", x) // Output: Value of x after modification: 20
}

func modifyPointerValue(p *int) {
	//In this example, we define a function called "modifyPointerValue" that takes a pointer to an integer as an argument.
	//Inside the function, we dereference the pointer using the "*" operator and assign it a new value of 20.
	*p = 20
}
