package main

import "fmt"

func printSlice[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func main() {
	intSlice := []int{1, 2, 3, 4, 5}
	stringSlice := []string{"a", "b", "c", "d", "e"}

	fmt.Println("Integer slice:")
	printSlice(intSlice)

	fmt.Println("String slice:")
	printSlice(stringSlice)
}
