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

func main() {
	rectangle := Rectangle{
		width:  25,
		height: 25,
	}

	fmt.Printf(
		"Width: %.2f, Height: %.2f, Area: %.2f\n",
		rectangle.width,
		rectangle.height,
		rectangle.Area(),
	)
}
