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

	// ========================================
	// GENERAL FORMATTING VERBS
	// ========================================
	// %v  : value in default format
	// %+v : structs with field names
	// %#v : Go-syntax representation
	// %T  : type of the value
	// %%  : literal percent sign

	//Example of general formatting:
	i := 1_5.5 //You can use underscores in numeric literals to improve readability. In this case, 1_5.5 is equivalent to 15.5. But not in the decimal part. For example, 1.5_5 is not valid.
	string := "My	name is John and I am 30 years old."

	fmt.Printf("Value: %v\n", i)             // Output: Value: 15.5
	fmt.Printf("Value: %#v\n", i)            // Output: Value: 15.5
	fmt.Printf("Type: %T\n", i)              // Output: Type: float64
	fmt.Printf("String: %s\n", string)       // Output: String: My name is John and I am 30 years old.
	fmt.Printf("String: %q\n", string)       // Output: String: "My name is John and I am 30 years old."
	fmt.Printf("String: %v\n", string)       // Output: String: "My name is John and I am 30 years old."
	fmt.Printf("String: %#v\n", string)      // Output: String: "My name is John and I am 30 years old."
	fmt.Printf("Literal percent sign: %%\n") // Output: Literal percent sign: %

	// ========================================
	// BOOLEANS
	// ========================================
	// %t  : true or false

	// ========================================
	// INTEGERS
	// ========================================
	// %b  : binary (base 2)
	// %c  : Unicode character represented by the value
	// %d  : decimal (base 10)
	// %o  : octal (base 8)
	// %O  : octal with 0o prefix
	// %q  : single-quoted character literal
	// %x  : hexadecimal (lowercase)
	// %X  : hexadecimal (uppercase)
	// %U  : Unicode format (U+1234)

	// ========================================
	// FLOATING-POINT & COMPLEX NUMBERS
	// ========================================
	// %b  : scientific notation, power of 2 exponent
	// %e  : scientific notation (e.g. 1.23e+10)
	// %E  : scientific notation (e.g. 1.23E+10)
	// %f  : decimal notation
	// %F  : same as %f
	// %g  : shortest between %e and %f
	// %G  : shortest between %E and %F

	// ========================================
	// STRINGS & BYTE SLICES
	// ========================================
	// %s  : raw string
	// %q  : double-quoted Go string
	// %x  : hexadecimal bytes (lowercase)
	// %X  : hexadecimal bytes (uppercase)

	// ========================================
	// POINTERS
	// ========================================
	// %p  : memory address in hexadecimal

	// ========================================
	// WIDTH & PRECISION MODIFIERS
	// ========================================
	// %8d     : width 8, right-aligned
	// %-8d    : width 8, left-aligned
	// %08d    : width 8, zero-padded
	// %.2f    : 2 decimal places
	// %8.2f   : width 8, precision 2
	// %*d     : width specified by argument
	// %.*f    : precision specified by argument

	// ========================================
	// EXAMPLES
	// ========================================
	// fmt.Printf("%v\n", value)
	// fmt.Printf("%+v\n", myStruct)
	// fmt.Printf("%#v\n", value)
	// fmt.Printf("%T\n", value)
	//
	// fmt.Printf("%t\n", true)
	//
	// fmt.Printf("%d\n", 255)
	// fmt.Printf("%b\n", 255)
	// fmt.Printf("%o\n", 255)
	// fmt.Printf("%x\n", 255)
	//
	// fmt.Printf("%f\n", 3.14159)
	// fmt.Printf("%.2f\n", 3.14159)
	//
	// fmt.Printf("%s\n", "hello")
	// fmt.Printf("%q\n", "hello")
	//
	// fmt.Printf("%p\n", &value)
}
