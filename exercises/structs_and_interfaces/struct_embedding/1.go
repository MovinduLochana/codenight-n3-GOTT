package main

import "fmt"

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// TODO: define NamedRectangle, embedding Rectangle, plus a Name string field

func main() {
	nr := NamedRectangle{Rectangle: Rectangle{Width: 3, Height: 4}, Name: "MyRect"}
	fmt.Println(nr.Name, nr.Area())
}