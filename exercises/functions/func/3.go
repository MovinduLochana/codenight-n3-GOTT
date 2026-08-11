package main

import "fmt"

func MinMax(nums []int) (min, max int) {
	// TODO: find and return the smallest and largest element in one pass
	return 0, 0
}

func main() {
	lo, hi := MinMax([]int{3, 1, 4, 1, 5})
	fmt.Println(lo, hi) // 1 5

	lo, hi = MinMax([]int{7})
	fmt.Println(lo, hi) // 7 7
}