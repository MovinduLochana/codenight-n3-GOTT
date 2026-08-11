package main

import "fmt"

func EvenSum(n int) int {
	// TODO: launch a goroutine that sends every even number 1..n
	// into a channel, then closes it.
	// In main, sum the values with `for v := range ch`.
	return 0
}

func main() {
	fmt.Println(EvenSum(10)) // 30
	fmt.Println(EvenSum(5))  // 6
	fmt.Println(EvenSum(0))  // 0
}
