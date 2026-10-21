package main

import (
	"fmt"
	"os"
	"strings"
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

	// ======================================================
	// PRINT TO CONSOLE
	// ======================================================

	// Print values without automatically adding a newline
	fmt.Print("Hello ")
	fmt.Print("World\n")

	// Print values separated by spaces and add a newline
	fmt.Println("Hello", "World")

	// Print using formatting verbs
	name = "Manuel"
	age = 36

	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Age: %d\n", age)

	// ======================================================
	// RETURN STRINGS INSTEAD OF PRINTING
	// ======================================================

	// Concatenate values into a string
	s1 := fmt.Sprint("Name: ", name)
	fmt.Println(s1)

	// Same as Sprint but adds spaces and a trailing newline
	s2 := fmt.Sprintln("Hello", "World")
	fmt.Print(s2)

	// Formatted string
	s3 := fmt.Sprintf("%s is %d years old", name, age)
	fmt.Println(s3)

	// ======================================================
	// WRITE TO AN io.Writer
	// ======================================================

	// Writes to stdout
	fmt.Fprint(os.Stdout, "Fprint example\n")

	// Writes to stdout and adds a newline
	fmt.Fprintln(os.Stdout, "Fprintln example")

	// Formatted write to stdout
	fmt.Fprintf(os.Stdout, "User: %s (%d)\n", name, age)

	// ======================================================
	// READ USER INPUT
	// ======================================================

	// Scan reads values separated by spaces
	var firstName string
	fmt.Print("Enter first name: ")
	fmt.Scan(&firstName)

	fmt.Println("You entered:", firstName)

	// Scanln reads until Enter is pressed
	var city string
	fmt.Print("Enter city: ")
	fmt.Scanln(&city)

	fmt.Println("City:", city)

	// Scanf reads using a format string
	var number int
	fmt.Print("Enter a number: ")
	fmt.Scanf("%d", &number)

	fmt.Println("Number:", number)

	// ======================================================
	// READ FROM STRINGS
	// ======================================================

	// Sscan parses values from a string
	var age1 int
	fmt.Sscan("36", &age1)
	fmt.Println(age1)

	// Sscanln parses until newline
	var age2 int
	fmt.Sscanln("40", &age2)
	fmt.Println(age2)

	// Sscanf parses according to a format
	var parsedAge int
	fmt.Sscanf("Age: 50", "Age: %d", &parsedAge)
	fmt.Println(parsedAge)

	// ======================================================
	// READ FROM io.Reader
	// ======================================================

	reader1 := strings.NewReader("100")
	var value1 int
	fmt.Fscan(reader1, &value1)
	fmt.Println(value1)

	reader2 := strings.NewReader("200")
	var value2 int
	fmt.Fscanln(reader2, &value2)
	fmt.Println(value2)

	reader3 := strings.NewReader("Age: 300")
	var value3 int
	fmt.Fscanf(reader3, "Age: %d", &value3)
	fmt.Println(value3)

	// ======================================================
	// APPEND TO []byte
	// ======================================================

	buffer := []byte("Hello")

	buffer = fmt.Append(buffer, " Manuel")
	fmt.Println(string(buffer))

	buffer = fmt.Appendln(buffer, " Appended line")
	fmt.Print(string(buffer))

	buffer = fmt.Appendf(buffer, " Age=%d", age)
	fmt.Println(string(buffer))

	// ======================================================
	// COMMON FORMAT VERBS
	// ======================================================

	// General
	fmt.Printf("%v\n", age) // default format
	fmt.Printf("%T\n", age) // type
	fmt.Printf("%%\n")      // literal %

	// String
	fmt.Printf("%s\n", "hello")
	fmt.Printf("%q\n", "hello")

	// Boolean
	fmt.Printf("%t\n", true)

	// Integers
	fmt.Printf("%d\n", 255) // decimal
	fmt.Printf("%b\n", 255) // binary
	fmt.Printf("%o\n", 255) // octal
	fmt.Printf("%x\n", 255) // hexadecimal lowercase
	fmt.Printf("%X\n", 255) // hexadecimal uppercase

	// Float
	fmt.Printf("%f\n", 123.456)
	fmt.Printf("%.2f\n", 123.456)
	fmt.Printf("%e\n", 123.456)
	fmt.Printf("%E\n", 123.456)
	fmt.Printf("%g\n", 123.456)
	fmt.Printf("%G\n", 123.456)

	// Character / Rune
	fmt.Printf("%c\n", 65) // A
	fmt.Printf("%U\n", 65) // U+0041

	// Pointer
	x := 10
	fmt.Printf("%p\n", &x)

	// ======================================================
	// WIDTH AND ALIGNMENT
	// ======================================================

	// Right align in 10 spaces
	fmt.Printf("|%10s|\n", "Go")

	// Left align in 10 spaces
	fmt.Printf("|%-10s|\n", "Go")

	// Integer padding
	fmt.Printf("|%5d|\n", 42)

	// Zero padding
	fmt.Printf("|%05d|\n", 42)

	// Float precision
	fmt.Printf("%.3f\n", 3.141592)

	// ======================================================
	// STRUCT EXAMPLES
	// ======================================================

	type User struct {
		Name string
		Age  int
	}

	user := User{
		Name: "Manuel",
		Age:  36,
	}

	fmt.Printf("%v\n", user)  // {Manuel 36}
	fmt.Printf("%+v\n", user) // {Name:Manuel Age:36}
	fmt.Printf("%#v\n", user) // Go syntax representation

	//Error formating functions
	errorAge := checkAge(16)
	fmt.Println(errorAge)
	errorAge = checkAge(22)
	fmt.Println(errorAge)
}

func checkAge(age int) error {
	if age < 18 {
		return fmt.Errorf("Age %d is too young to die", age)
	}
	return nil
}
