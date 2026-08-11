package main

import "fmt"

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// TODO: define WeightedRect embedding Rectangle, plus a Weight float64 field
// TODO: give WeightedRect its OWN Area() method that returns
//   promoted area * Weight  (nr.Rectangle.Area() * nr.Weight)

func main() {
	wr := WeightedRect{Rectangle: Rectangle{Width: 3, Height: 4}, Weight: 2}
	fmt.Println(wr.Area())          // 24
	fmt.Println(wr.Rectangle.Area()) // 12
}