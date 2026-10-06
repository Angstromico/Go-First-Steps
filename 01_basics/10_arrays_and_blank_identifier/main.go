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

	//If you assing an array to another array, it will create a copy of the array and not a reference to the original array.
	arr3 := arr
	arr3[0] = 99
	fmt.Println(arr)  // Original array remains unchanged
	fmt.Println(arr3) // Modified copy of the array

	//Loop in array with for loop:

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
