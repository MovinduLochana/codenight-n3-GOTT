package main

import (
	"fmt"
	"sync"
)

func ConcurrentSum(nums []int) int {
	// TODO: for each number in nums, launch a goroutine that adds it
	// to a shared `total` guarded by a sync.Mutex.
	// Use a sync.WaitGroup; return the final total.
	return 0
}

func main() {
	fmt.Println(ConcurrentSum([]int{1, 2, 3, 4})) // 10
	fmt.Println(ConcurrentSum([]int{}))           // 0
}
