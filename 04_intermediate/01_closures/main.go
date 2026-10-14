package main

import "fmt"

func main() {
	//Closures are functions that are defined inside another function and can access the variables of the outer function.
	//In this example, we define a function called "outer" that takes an integer parameter "x" and returns a function that takes an integer parameter "y".
	//The returned function adds "x" and "y" together and returns the result.
	//We then call the "outer" function with the value 5 and assign the returned function to a variable called "add5".
	//Finally, we call the "add5" function with the value 10 and print the result.
	outer := func(x int) func(y int) int {
		return func(y int) int {
			return x + y
		}
	}
	add5 := outer(5)
	fmt.Println(add5(10)) // Output: 15

	sequence := adder()
	fmt.Println(sequence()) // Output: 1
	fmt.Println(sequence()) // Output: 2
	fmt.Println(sequence()) // Output: 3
	fmt.Println(sequence()) // Output: 4
}

func adder() func() int {
	i := 0
	fmt.Println("i is", i)
	return func() int {
		i++
		println("i is", i)
		return i
	}
}
