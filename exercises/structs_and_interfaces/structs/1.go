package main

import "fmt"

// TODO: define the Rectangle struct with Width and Height (float64)

func Area(r Rectangle) float64 {
	// TODO: return r.Width * r.Height
	return 0
}

func main() {
	r := Rectangle{Width: 3, Height: 4}
	fmt.Println(Area(r))
}