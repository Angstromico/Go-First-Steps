package main

import (
	"fmt"
	"maps"
)

func main() {
	// Create a map
	mapExample := map[string]int{
		"one": 1,
		"two": 2,
	}
	fmt.Println(mapExample)

	// Access a value
	fmt.Println("Value of 'one':", mapExample["one"])

	// Add a new key-value pair
	mapExample["three"] = 3
	fmt.Println(mapExample)

	// Delete a key-value pair
	delete(mapExample, "two")
	fmt.Println(mapExample)

	//Init an empty map
	emptyMap := make(map[string]int)
	fmt.Println(emptyMap)

	//Loop in map
	for key, value := range mapExample {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}

	//Check if a key exists
	if value, exists := mapExample["one"]; exists {
		fmt.Println("Key 'one' exists with value:", value)
	} else {
		fmt.Println("Key 'one' does not exist")
	}

	//What happen if you consult a key that does not exist
	fmt.Println("Value of 'nonexistent':", mapExample["nonexistent"]) //Output: 0

	//Get the key and the value in a separate way
	keys := make([]string, 0, len(mapExample))
	for key := range mapExample {
		keys = append(keys, key)
	}
	fmt.Println("Keys:", keys)

	values := make([]int, 0, len(mapExample))
	for _, value := range mapExample {
		values = append(values, value)
	}
	fmt.Println("Values:", values)

	//Clear the map
	clear(mapExample)
	fmt.Println(mapExample)

	//Check if an value is asociate to some key
	_, exists := mapExample["one"]
	if exists {
		fmt.Println("Key 'one' exists")
	} else {
		fmt.Println("Key 'one' does not exist")
	}

	// maps.Equal compares two maps for equality
	// Example:
	map1 := map[string]int{"a": 1, "b": 2}
	map2 := map[string]int{"a": 1, "b": 2}
	fmt.Println(maps.Equal(map1, map2)) // Output: true

	//Any uninitialized map will have a nil value, and accessing it will panic.
	// Example:
	var nilMap map[string]int
	fmt.Println(nilMap == nil) // This will cause a panic
	//In order to app a new key to a nil map, you need to initialize it first.
	// Example:
	nilMap = make(map[string]int)
	nilMap["key"] = 1
	fmt.Println(nilMap)
}
