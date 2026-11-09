package main

import "fmt"

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) Pop() (T, bool) {
	var zero T

	if len(s.elements) == 0 {
		return zero, false
	}

	lastIndex := len(s.elements) - 1
	element := s.elements[lastIndex]
	s.elements = s.elements[:lastIndex]

	return element, true
}

func printSlice[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func swap[T any](a, b *T) {
	*a, *b = *b, *a
}

func main() {
	stack := Stack[int]{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	element, ok := stack.Pop()
	fmt.Println("Popped element:", element, "ok:", ok)
	element, ok = stack.Pop()
	fmt.Println("Popped element:", element, "ok:", ok)
	element, ok = stack.Pop()
	fmt.Println("Popped element:", element, "ok:", ok)

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
