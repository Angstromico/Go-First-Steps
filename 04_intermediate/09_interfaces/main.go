package main

import (
	"fmt"
	"math"
)

type Geometry interface {
	area() float64
	perim() float64
}

type Rect struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rect) area() float64 {
	return r.height * r.width
}
func (a Circle) area() float64 {
	return math.Pi * a.radius * a.radius
}

func (p Circle) perim() float64 {
	return 2 * math.Pi * p.radius
}
func (p Rect) perim() float64 {
	return 2 * (p.width + p.height)
}

func measure(g Geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	rectangle := Rect{
		width:  8,
		height: 9,
	}

	circle := Circle{
		radius: 5,
	}

	measure(rectangle)
	measure(circle)
}
