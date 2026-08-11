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

// TODO: define Circle struct with field Radius float64
// TODO: give Circle an Area() float64 method (use 3.14159 for pi)

func TotalArea(shapes []Shape) float64 {
	// TODO: sum the Area() of every shape in the slice
	return 0
}

func main() {
	shapes := []Shape{
		Rectangle{Width: 2, Height: 3},
		Circle{Radius: 1},
	}
	fmt.Println(TotalArea(shapes))
}