package main

import (
	"fmt"
	"os"
)

func main() {
	//The fmt package in Go provides functions for formatting and printing output. It is part of the standard library and is widely used for displaying information to the console or writing formatted data to files.
	//In this example, we use the "fmt.Println" function to print a simple message to the console. The "Println" function automatically adds a newline character at the end of the output.
	fmt.Println("Hello, World!") // Output: Hello, World!

	//The "fmt.Sprintf" function is used to format a string without printing it to the console. It returns the formatted string, which can be stored in a variable or used in other ways.
	name := "Alice"
	age := 25
	message := fmt.Sprintf("My name is %s and I am %d years old.", name, age)
	fmt.Println(message) // Output: My name is Alice and I am 25 years old.

	//The "fmt.Errorf" function is used to create formatted error messages. It returns an error value that can be used in error handling.
	err := fmt.Errorf("An error occurred: %s", "something went wrong")
	fmt.Println(err) // Output: An error occurred: something went wrong

	//The "fmt.Fprint" function is used to write formatted output to a specified writer, such as a file or a network connection. In this example, we use "os.Stdout" as the writer to print to the console.
	//Note: To use "os.Stdout", you need to import the "os" package.
	fmt.Fprint(os.Stdout, "This is a message written to the console using Fprint.\n")

	//The "fmt.Fprintln" function is similar to "fmt.Fprint", but it automatically adds a newline character at the end of the output.
	fmt.Fprintln(os.Stdout, "This is a message written to the console using Fprintln.")

	// The "fmt.Printf" function is used to format and print output to the console. It allows you to include formatting verbs in the string, which are replaced with the corresponding values.
	//In this example, we use the "%s" verb to format a string and the "%d" verb to format an integer.
	fmt.Printf("My name is %s and I am %d years old.\n", name, age) // Output: My name is Alice and I am 25 years old.
}
