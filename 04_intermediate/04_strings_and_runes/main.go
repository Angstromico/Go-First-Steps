package main

import "fmt"

func main() {
	//In this example, we define a string variable called "s" and assign it the value "Hello, World!".
	//We then use the "len" function to get the length of the string and print it to the console.
	s := "Hello, World!"
	fmt.Println("Length of string:", len(s)) // Output: Length of string: 13

	//Runes are a data type in Go that represent Unicode code points. They are similar to characters in other programming languages.
	//In this example, we define a rune variable called "r" and assign it the value 'A'.
	//We then use the "fmt.Printf" function to print the Unicode code point of the rune to the console.
	r := 'A'
	fmt.Printf("Unicode code point of rune: %U\n", r) // Output: Unicode code point of rune: U+0041

	//In strings you can use backticks to create raw string literals that can span multiple lines and include special characters without the need for escape sequences.
	//In this example, we define a raw string literal called "rawString" using backticks.
	//We then print the raw string literal to the console.
	rawString := `This is a raw string literal.
	It can span multiple lines and include special characters like \n and \t without the need for escape sequences.`
	fmt.Println("Raw string literal:")
	fmt.Println(rawString)
}
