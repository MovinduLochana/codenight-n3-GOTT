package main

import "fmt"

func InsertAt(s []int, i, v int) []int {
	// TODO: return a new slice with v inserted at index i, order preserved
	// hint: append a zero value to grow s, then re-slice and shift
	return nil
}

func main() {
	fmt.Println(InsertAt([]int{10, 20, 30}, 1, 99)) // [10 99 20 30]
	fmt.Println(InsertAt([]int{1, 2}, 0, 0))        // [0 1 2]
}