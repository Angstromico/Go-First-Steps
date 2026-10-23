package main

import "fmt"

type Rectangle struct {
	width  float64
	height float64
}

// Method with value reciever
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// Method with pointer reciever
func (r *Rectangle) Scale(factor float64) {
	r.height *= factor
	r.width *= factor
}

func main() {
	rectangle := Rectangle{
		width:  25,
		height: 25,
	}

	rectangle.Scale(5)

	fmt.Printf(
		"Width: %.2f, Height: %.2f, Area: %.2f\n",
		rectangle.width,
		rectangle.height,
		rectangle.Area(),
	)

	number := MyInt(-5)
	isPositive := number.IsPositive()

	fmt.Printf("The number %d is positive: %t\n", number, isPositive)
}

type MyInt int

// Method on user-defined type
func (m MyInt) IsPositive() bool {
	return m > 0
}
