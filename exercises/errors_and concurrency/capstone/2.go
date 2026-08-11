package main

import (
	"fmt"
	"sync"
)

func CheckoutTotals() int {
	// Buffered so the producer can send all 3 values and close without
	// blocking — the unbuffered variant would deadlock the wg.Wait() above.
	totals := make(chan int, 3)

	var wg sync.WaitGroup
	wg.Add(1)
	// TODO: launch a goroutine (inside, defer wg.Done()) that sends
	// 10, 20, 30 into `totals`, then close(totals)

	wg.Wait()

	// TODO: sum the values with `for v := range totals` and return the total
	return 0
}

func main() {
	fmt.Println(CheckoutTotals()) // 60
}
