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
}
