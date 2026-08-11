package main

import "fmt"

type Rectangle struct {
	Width, Height float64
}

// TODO: define Scale(factor float64) on *Rectangle

func main() {
	r := Rectangle{Width: 3, Height: 4}
	r.Scale(2)
	fmt.Println(r.Width, r.Height)
}