package main

import "fmt"

func main() {
	//Formatting verbs are used in Go to format strings and other data types for output. They are used with the "fmt" package and are represented by a percent sign followed by a letter that indicates the type of data being formatted.
	//In this example, we define a string variable called "name" and assign it the value "John".
	//We then use the "fmt.Printf" function to format the string using the "%s" verb, which indicates that we are formatting a string.
	//We also use the "%d" verb to format an integer variable called "age" and print it to the console.
	name := "John"
	age := 30
	fmt.Printf("My name is %s and I am %d years old.\n", name, age) // Output: My name is John and I am 30 years old.
}
