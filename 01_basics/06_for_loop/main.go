package main

func main() {
	//	For loop
	for i := 0; i < 5; i++ {
		println("Iteration:", i)
	}
	for i := 0; i < len([]string{"Go", "is", "fun"}); i++ {
		println("Word:", []string{"Go", "is", "fun"}[i])
	}
	//	Range loop
	words := []string{"Go", "is", "fun"}
	for _, word := range words {
		println("Word:", word)
	}
	// Odd or even numbers loop
	const limit = 10

	for i := range make([]int, limit) {
		if i%2 == 0 {
			println(i, "is even")
		} else {
			println(i, "is odd")
		}
	}

	rows := 5

	//	Nested loops to print a pattern
	for i := 0; i < rows; i++ {
		for j := 0; j <= i; j++ {
			print("* ")
		}
		println()
	}

	//Outer loop
	for i := 0; i < 10; i++ {
		//Inner loop for spaces before stars
		for j := 0; j < 10-i-1; j++ {
			print(" ")
		}
		//Inner loop for stars
		for j := 0; j < 2*i+1; j++ {
			print("*")
		}
		println() //Moving to the next line after each row of stars
	}
	for value := range make([]int, 5) {
		println("Value from range loop:", value)
	}
	//While loop	using for
	counter := 0
	for counter < 5 {
		println("Counter:", counter)
		counter++
	}
	sum := 0
	for {
		sum += 10
		print("Sum:", sum)
		print("\n")
		if sum >= 50 {
			break
		}
	}
}
