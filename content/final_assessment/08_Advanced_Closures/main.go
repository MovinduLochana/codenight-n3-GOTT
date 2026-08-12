package main

import "fmt"

// TODO: Implement MakeFilter(limit int) func([]int) []int

func main() {
	nums := []int{10, 5, 20, 15, 30, 2}

	filterAbove10 := MakeFilter(10)
	filterAbove20 := MakeFilter(20)

	fmt.Println("Above 10:", filterAbove10(nums))
	fmt.Println("Above 20:", filterAbove20(nums))
}
