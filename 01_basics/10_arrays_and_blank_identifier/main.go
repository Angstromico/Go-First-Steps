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

	//For loop range and value in same message:

	for i, v := range arr {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}

	//Using the blank identifier to ignore values returned by a function
	_, b := someFunction()
	fmt.Println(b)

	//Arrays with same content are not equal
	arr4 := [3]int{1, 2, 3}
	fmt.Println(arr == arr4) // false
	arr5 := [3]int{1, 2, 4}
	fmt.Println(arr == arr5) // false
}

func someFunction() (int, int) {
	return 1, 2
}
