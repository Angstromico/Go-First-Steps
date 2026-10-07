package main

import "fmt"

func main() {
	// Example of using range with a slice
	slice := []int{1, 2, 3, 4, 5}
	for index, value := range slice {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// Example of using range with a map
	mapExample := map[string]int{
		"one": 1,
		"two": 2,
	}
	for key, value := range mapExample {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
}
