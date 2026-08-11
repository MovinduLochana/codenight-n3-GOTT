package main

import "fmt"

func MakeCounter() func() int {
	// TODO: return a closure that increments and returns a count, starting at 1
	return nil
}

func main() {
	counter := MakeCounter()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
}