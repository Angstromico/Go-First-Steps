package main

import "fmt"

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
}
