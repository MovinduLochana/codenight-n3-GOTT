package main

import "fmt"

func SquaresPipeline(k int) []int {
	// TODO: build a two-stage pipeline with channels:
	//   1. producer goroutine sends 1..k into `nums`
	//   2. transformer goroutine reads nums, squares, sends into `squares`
	//   3. main drains squares with for range, appending to the result
	return nil
}

func main() {
	fmt.Println(SquaresPipeline(4)) // [1 4 9 16]
	fmt.Println(SquaresPipeline(0)) // []
}
