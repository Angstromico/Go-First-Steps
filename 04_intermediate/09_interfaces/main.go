package main

import (
	"fmt"
	"math"
)

type Geometry interface {
	area() float64
	perim() float64
}

type Animal interface {
	Speak() string
}

type Rect struct {
	width, height float64
}

type Circle struct {
	radius float64
}

type Dog struct {
	Name string
}

type Cat struct {
	Name string
}

func (r Rect) area() float64 {
	return r.height * r.width
}
func (a Circle) area() float64 {
	return math.Pi * math.Pow(a.radius, 2)
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

func (d Dog) Speak() string {
	return "Woof!"
}

func (c Cat) Speak() string {
	return "Meow!"
}

func makeAnimalSpeak(a Animal) {
	fmt.Printf("%T says: %s\n", a, a.Speak())
}

func main() {
	rectangle := Rect{
		width:  8,
		height: 9,
	}

	dog := Dog{Name: "Rex"}
	cat := Cat{Name: "Mittens"}

	circle := Circle{
		radius: 5,
	}

	measure(rectangle)
	measure(circle)

	makeAnimalSpeak(dog)
	makeAnimalSpeak(cat)
}
