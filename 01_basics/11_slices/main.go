package main

import "fmt"

func main() {
	// Create a slice
	slice := []int{1, 2, 3}
	fmt.Println(slice)

	// Append to a slice
	slice = append(slice, 4)
	fmt.Println(slice)

	// Slice of slices
	sliceOfSlices := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	for i, row := range sliceOfSlices {
		for j, val := range row {
			fmt.Printf("sliceOfSlices[%d][%d] = %d\n", i, j, val)
		}
	}

	//Assing an array to a slice
	arr := [3]int{7, 8, 9}
	sliceFromArr := arr[:]
	limitArrayFromSlice := sliceFromArr[:2] //[7,8,9]
	evenMoreLimitArrayFromSlice := sliceFromArr[:1]

	fmt.Println(sliceFromArr)
	fmt.Println(limitArrayFromSlice)
	fmt.Println(evenMoreLimitArrayFromSlice)
}
