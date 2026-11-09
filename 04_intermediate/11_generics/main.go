package main

import "fmt"

func printSlice[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func swap[T any](a, b *T) {
	*a, *b = *b, *a
}

func main() {
	intSlice := []int{1, 2, 3, 4, 5}
	stringSlice := []string{"a", "b", "c", "d", "e"}

	fmt.Println("Integer slice:")
	printSlice(intSlice)

	fmt.Println("String slice:")
	printSlice(stringSlice)

	fmt.Println("Swapping two integers:")
	x, y := 10, 20
	swap(&x, &y)
	fmt.Println("x:", x, "y:", y)

	fmt.Println("Swapping two strings:")
	s1, s2 := "hello", "world"
	swap(&s1, &s2)
	fmt.Println("s1:", s1, "s2:", s2)
}
