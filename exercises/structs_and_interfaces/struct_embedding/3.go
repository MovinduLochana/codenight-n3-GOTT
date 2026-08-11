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

// TODO: define Square embedding Rectangle, plus a Side float64 field
// TODO: give Square its OWN Area() method returning Side * Side,
//   shadowing the promoted Rectangle.Area()

func TotalArea(shapes []Shape) float64 {
	// TODO: sum the Area() of every shape
	return 0
}

func main() {
	shapes := []Shape{
		Square{Rectangle: Rectangle{Width: 3, Height: 4}, Side: 5},
	}
	fmt.Println(TotalArea(shapes))
}