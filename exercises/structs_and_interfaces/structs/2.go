package main

import "fmt"

// TODO: define the Box struct with Length, Width, Height (all float64)

func Volume(b Box) float64 {
	// TODO: return b.Length * b.Width * b.Height
	return 0
}

func main() {
	fmt.Println(Volume(Box{Length: 2, Width: 3, Height: 4}))
}