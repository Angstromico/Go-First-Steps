package main

import (
	"fmt"
	"slices"
)

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

	//Nil slice
	var nilSlice []int
	fmt.Println(nilSlice)

	//Compare Slices
	slice1 := []int{1, 2, 3}
	slice2 := []int{1, 2, 3}
	slice3 := []int{1, 2, 4}

	if slices.Equal(slice1, slice2) {
		fmt.Println("slice1 and slice2 are equal")
	}
	if slices.Equal(slice1, slice3) {
		fmt.Println("slice1 and slice3 are equal")
	}

	//slice[low:hight]
	slice4 := []int{1, 2, 3, 4, 5}
	subSlice := slice4[1:3] // [2, 3]
	fmt.Println(slice4)
	fmt.Println(subSlice)

	//Capacity of a slice
	fmt.Println("Capacity of slice4:", cap(slice4))
	fmt.Println("Length of slice4:", len(slice4))

	fmt.Println(len(subSlice)) // 2
	fmt.Println(cap(subSlice)) // 4
}
