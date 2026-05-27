package main

func main() {
	// Arithmetic operators
	a := 10
	b := 5
	c := a + b // Addition
	d := a - b // Subtraction
	e := a * b // Multiplication
	f := a / b // Division
	g := a % b // Modulus

	println("Addition:", c)
	println("Subtraction:", d)
	println("Multiplication:", e)
	println("Division:", f)
	println("Modulus:", g)

	//	Compound assignment operators
	a += 2 // a = a + 2
	println("After compound addition (a += 2):", a)
	a -= 2 // a = a - 2
	println("After compound subtraction (a -= 2):", a)
	a *= 2 // a = a * 2
	println("After compound multiplication (a *= 2):", a)
	a /= 2 // a = a / 2
	println("After compound division (a /= 2):", a)

	//	Increment and decrement operators
	a++ // a = a + 1
	println("After increment (a++):", a)
	a-- // a = a - 1
	println("After decrement (a--):", a)

	//	Operator precedence
	result := a + b*2 // Multiplication has higher precedence than addition
	println("Result with operator precedence:", result)
	result = (a + b) * 2 // Parentheses change the order of evaluation
	println("Result with parentheses:", result)
}
