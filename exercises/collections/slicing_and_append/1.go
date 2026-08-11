package main

import "fmt"

func Cut(s []int, lo, hi int) []int {
	// TODO: return the sub-slice s[lo:hi]
	return nil
}

func main() {
	fmt.Println(Cut([]int{10, 20, 30, 40, 50}, 1, 3)) // [20 30]
	fmt.Println(Cut([]int{5, 6, 7}, 0, 2))            // [5 6]
}