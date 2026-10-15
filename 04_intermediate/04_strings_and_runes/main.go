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

	//Normal string literals are enclosed in double quotes and can include escape sequences to represent special characters.
	//In this example, we define a normal string literal called "normalString" using double quotes.
	//We then print the normal string literal to the console.
	normalString := "This is a normal string literal.\nIt can include escape sequences like \\n and \\t."
	fmt.Println("Normal string literal:")
	fmt.Println(normalString)

	//Accesing a character in a string by its index:
	//In this example, we define a string variable called "str" and assign it the value "Hello, World!".
	//We then access the character at index 7 of the string using the indexing operator and print it to the console.
	str := "Hello, World!"
	fmt.Println("Character at index 7:", string(str[7])) // Output: Character at index 7: W

	//String concatenation is the process of combining two or more strings into a single string.
	//In this example, we define two string variables called "str1" and "str2" and assign them the values "Hello, " and "World!" respectively.
	//We then concatenate the two strings using the "+" operator and print the result to the console.
	str1 := "Hello, "
	str2 := "World!"
	concatenatedString := str1 + str2
	fmt.Println("Concatenated string:", concatenatedString) // Output: Concatenated string: Hello, World!

	//Compare two strings using the "==" operator. In this example, we define two string variables called "str3" and "str4" and assign them the values "Hello" and "Hello" respectively. We then compare the two strings using the "==" operator and print the result to the console.
	str3 := "Hello"
	str4 := "Hello"
	fmt.Println("Are str3 and str4 equal?", str3 == str4) // Output: Are str3 and str4 equal? true
	//Compare two strings using the "!=" operator. In this example, we define two string variables called "str5" and "str6" and assign them the values "Hello" and "World" respectively. We then compare the two strings using the "!=" operator and print the result to the console.
	str5 := "Hello"
	str6 := "World"
	fmt.Println("Are str5 and str6 not equal?", str5 != str6) // Output: Are str5 and str6 not equal? true
	//Compare two strings using the ">" operator. In this example, we define two string variables called "str7" and "str8" and assign them the values "Hello" and "World" respectively. We then compare the two strings using the ">" operator and print the result to the console.
	str7 := "Hello"
	str8 := "World"
	fmt.Println("Is str7 greater than str8?", str7 > str8)
	//A simple lowercase can made a string greater than another string. In this example, we define two string variables called "str9" and "str10" and assign them the values "hello" and "Hello" respectively. We then compare the two strings using the ">" operator and print the result to the console.
	str9 := "hello"
	str10 := "Hello"
	fmt.Println("Is str9 greater than str10?", str9 > str10) // Output: Is str9 greater than str10? true
}
