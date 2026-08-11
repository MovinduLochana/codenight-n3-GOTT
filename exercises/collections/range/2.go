package main

import "fmt"

func Contains(nums []int, target int) bool {
	// TODO: use range (ignore index) to find target, return early
	return false
}

func main() {
	fmt.Println(Contains([]int{1, 2, 3}, 2)) // true
	fmt.Println(Contains([]int{1, 2, 3}, 9)) // false
	fmt.Println(Contains([]int{}, 0))        // false
}