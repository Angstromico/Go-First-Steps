package main

import "fmt"

func main() {
	arr := [3]int{1, 2, 3}
	for _, v := range arr {
		fmt.Println(v)
	}
}
