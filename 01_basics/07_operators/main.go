package main

import "fmt"

func main() {
	// Arithmetic operators
	a := 10
	b := 3
	fmt.Println("a + b =", a+b)
	fmt.Println("a - b =", a-b)
	fmt.Println("a * b =", a*b)
	fmt.Println("a / b =", a/b)
	fmt.Println("a % b =", a%b)

	// Comparison operators
	fmt.Println("a == b =", a == b)
	fmt.Println("a != b =", a != b)
	fmt.Println("a > b =", a > b)
	fmt.Println("a < b =", a < b)
	fmt.Println("a >= b =", a >= b)
	fmt.Println("a <= b =", a <= b)

	// Logical operators
	fmt.Println("a > 5 && b < 5 =", a > 5 && b < 5)
	fmt.Println("a > 5 || b < 2 =", a > 5 || b < 2)
	fmt.Println("! (a > 5) =", !(a > 5))

	// Bitwise operators
	fmt.Println("a & b =", a&b)
	fmt.Println("a | b =", a|b)
	fmt.Println("a ^ b =", a^b)
	fmt.Println("~a =", ^a)
}
