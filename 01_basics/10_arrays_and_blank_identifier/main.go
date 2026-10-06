package main

import "fmt"

func main() {
	arr := [3]int{1, 2, 3}
	//Update last element of the array
	arr[2] = 10
	for _, v := range arr {
		fmt.Println(v)
	}

	// Make an array of an undeterminate number of values
	arr2 := [...]int{4, 5, 6}
	for _, v := range arr2 {
		fmt.Println(v)
	}
}
