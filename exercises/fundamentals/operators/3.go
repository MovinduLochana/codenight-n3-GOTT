package main

import "fmt"

func IsBetween(n, low, high int) bool {
	// TODO: return true if n is in the inclusive range low..high
	return false
}

func main() {
	fmt.Println(IsBetween(5, 1, 10))  // should print true
	fmt.Println(IsBetween(0, 1, 10))  // should print false
	fmt.Println(IsBetween(10, 1, 10)) // should print true
}
