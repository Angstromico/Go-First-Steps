package main

import "fmt"

func main() {
	//Recursion is a programming technique where a function calls itself in order to solve a problem.
	//In this example, we define a function called "factorial" that takes an integer parameter "n" and returns the factorial of "n".
	//The factorial of a number is the product of all positive integers less than or equal to that number.
	//We use recursion to calculate the factorial by calling the "factorial" function with the value of "n-1" until we reach the base case of "n" being equal to 0.
	//When "n" is equal to 0, we return 1, which is the base case for the factorial function.
	//Finally, we call the "factorial" function with the value 5 and print the result.
	fmt.Println(factorial(5)) // Output: 120

	//In this example, we define a function called "fibonacci" that takes an integer parameter "n" and returns the nth number in the Fibonacci sequence.
	//The Fibonacci sequence is a series of numbers where each number is the sum of the two preceding ones, starting from 0 and 1.
	//We use recursion to calculate the nth number in the Fibonacci sequence by calling the "fibonacci" function with the values of "n-1" and "n-2" until we reach the base case of "n" being less than or equal to 1.
	//When "n" is less than or equal to 1, we return "n", which is the base case for the Fibonacci function.
	//Finally, we call the "fibonacci" function with the value 10 and print the result.
	fmt.Println(fibonacci(10)) // Output: 55
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
