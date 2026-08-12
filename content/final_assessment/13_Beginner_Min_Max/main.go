package main

import (
	"errors"
	"fmt"
)

// TODO: Implement FindMinMax(nums []int) (int, int, error)
// It should find the min and max values in the slice.
// If the slice is empty, return 0, 0, and an error with message "empty slice".

func main() {
	nums := []int{12, 45, 2, 8, 99, 14}
	min, max, err := FindMinMax(nums)
	if err == nil {
		fmt.Printf("Min: %d, Max: %d\n", min, max)
	}
}
