package main

import "fmt"

func main() {
	a := 10
	b := 3
	switch {
	case a == b:
		fmt.Println("a is equal to b")
	case a > b:
		fmt.Println("a is greater than b")
	default:
		fmt.Println("a is less than b")
	}
}
