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

	//Overflow and underflow
	var maxInt int = 1<<63 - 1
	println("Maximum int value:", maxInt)
	var minInt int = -1 << 63
	println("Minimum int value:", minInt)
	var overflowInt int = maxInt + 1
	println("Overflowed int value (maxInt + 1):", overflowInt)
	var underflowInt int = minInt - 1
	println("Underflowed int value (minInt - 1):", underflowInt)
	var maxFloat float64 = 1.7976931348623157e+308
	println("Maximum float value:", maxFloat)
	var underflowFloat float64 = 5e-324
	println("Minimum float value:", underflowFloat)
	maxInt = maxInt + 1
	println("Overflowed int value (maxInt + 1):", maxInt)
	underflowInt = minInt - 1
	println("Underflowed int value (minInt - 1):", underflowInt)
	println("Max int + 1 =", int64(maxInt)+1)
	println("Min int - 1 =", int64(minInt)-1)
}
