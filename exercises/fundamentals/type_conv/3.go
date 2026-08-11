package main

import "fmt"

func ToPixels(inches, dpi float64) int {
	// TODO: multiply inches * dpi in float64 first,
	// then convert the result to int with int(...)
	return 0
}

func main() {
	fmt.Println(ToPixels(1.5, 300)) // should print 450
	fmt.Println(ToPixels(0.5, 72))  // should print 36
}
