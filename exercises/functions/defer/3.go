package main

import "fmt"

func LoopCleanup(n int) (nums []int) {
	// TODO: for i := 0; i < n; i++ {
	//   defer appending i to the named return value `nums`
	// }
	// Deferred calls run after the loop in LIFO order,
	// so the result comes back reversed: n-1 ... 0
	return
}

func main() {
	fmt.Println(LoopCleanup(3)) // [2 1 0]
	fmt.Println(LoopCleanup(1)) // [0]
}