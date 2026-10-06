package main

import "fmt"

func main() {
	arr := [3]int{1, 2, 3}
	//Update last element of the array
	arr[2] = 10
	for _, v := range arr {
		fmt.Println(v)
	}
}
