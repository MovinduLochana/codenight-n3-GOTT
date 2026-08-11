package main

import "fmt"

// TODO: define interface Perimeterer with method Perimeter() float64

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Perimeter() float64 {
	// TODO: return 2 * (Width + Height)
	return 0
}

type Circle struct {
	Radius float64
}

func (c Circle) Perimeter() float64 {
	// TODO: return 2 * 3.14159 * Radius
	return 0
}

func TotalPerimeter(shapes []Perimeterer) float64 {
	// TODO: sum every shape's Perimeter()
	return 0
}

func main() {
	shapes := []Perimeterer{
		Rectangle{Width: 2, Height: 3},
		Circle{Radius: 1},
	}
	fmt.Println(TotalPerimeter(shapes))
}