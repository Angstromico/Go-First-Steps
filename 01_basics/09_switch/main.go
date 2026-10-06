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

	//Using fallthrough
	switch a {
	case 10:
		fmt.Println("a is 10")
		fallthrough
	case 3:
		fmt.Println("a is 3")
	default:
		fmt.Println("a is something else")
	}

	day := "Monday"
	switch day {
	case "Monday", "Tuesday":
		fmt.Println("It's a weekday")
	default:
		fmt.Println("It's a weekend")
	}

	checkType(42)
	checkType("hello")
	checkType(3.14)
}

func checkType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("v is an int")
	case string:
		fmt.Println("v is a string")
	default:
		fmt.Println("v is of some other type")
	}
}
