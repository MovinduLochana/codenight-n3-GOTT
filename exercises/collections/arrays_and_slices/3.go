package main

import "fmt"

func Tail(s []int, n int) []int {
	// TODO: return an independent copy of the last n elements
	// use make + copy; if n >= len(s) copy the whole slice
	return nil
}

func main() {
	fmt.Println(Tail([]int{1, 2, 3, 4, 5}, 2)) // [4 5]
	fmt.Println(Tail([]int{1, 2}, 5))          // [1 2]
	fmt.Println(Tail([]int{1, 2, 3}, 0))       // []
}