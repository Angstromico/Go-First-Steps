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
}
