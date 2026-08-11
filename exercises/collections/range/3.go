package main

import "fmt"

func IndexAll(nums []int, target int) []int {
	// TODO: use range with index to collect every position of target
	return nil
}

func main() {
	fmt.Println(IndexAll([]int{1, 2, 1, 4, 1}, 1)) // [0 2 4]
	fmt.Println(IndexAll([]int{1, 2, 3}, 9))       // []
}