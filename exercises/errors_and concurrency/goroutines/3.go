package main

import (
	"fmt"
	"sync"
)

func Squares(nums []int) []int {
	// TODO: launch one goroutine per number; each squares its value and
	// stores it into a shared result slice guarded by a sync.Mutex.
	// Wait for all goroutines, then SORT the result before returning.
	return nil
}

func main() {
	fmt.Println(Squares([]int{4, 1, 3})) // [1 9 16]
	fmt.Println(Squares([]int{}))        // []
}
