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

	//	Bitwise operators
	x := 5                               // In binary: 0101
	y := 3                               // In binary: 0011
	println("Bitwise AND (x & y):", x&y) // 1 (0001)
	println("Bitwise OR (x | y):", x|y)  // 7 (0111)

	// Interger on float	operations
	var f1 float64 = 3.5
	var i1 int = 2
	resultFloat := f1 + float64(i1)
	println("Result of integer to float operation:", resultFloat) //This is a float64 result
	const p float64 = 10 / 3
	println("Result of integer division assigned to float constant:", p) //This will be 3, because the division is done as integer division before being assigned to the float constant
	const q float64 = 10.0 / 3
	println("Result of float division assigned to float constant:", q) //This will be approximately 3.3333333333333335, because the division is done as float division
}
