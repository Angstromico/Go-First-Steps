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

	//Evaluating a case:
	switch a {
	case 10:
		fmt.Println("a is 10")
	case 3:
		fmt.Println("a is 3")
	default:
		fmt.Println("a is something else")
	}
}
