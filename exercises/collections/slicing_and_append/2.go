package main

import "fmt"

func RemoveAt(s []int, i int) []int {
	// TODO: return a new slice with the element at index i removed
	return s
}

func main() {
	fmt.Println(RemoveAt([]int{10, 20, 30}, 1)) // [10 30]
	fmt.Println(RemoveAt([]int{5, 6, 7, 8}, 0)) // [6 7 8]
}