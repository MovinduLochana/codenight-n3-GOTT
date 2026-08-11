package main

import "fmt"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

func Biggest(shapes []Shape) float64 {
	// TODO: return the largest Area() among the shapes
	return 0
}

func main() {
	shapes := []Shape{
		Rectangle{Width: 2, Height: 3},
		Circle{Radius: 2},
	}
	fmt.Println(Biggest(shapes))
}