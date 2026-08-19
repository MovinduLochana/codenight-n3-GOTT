package main

import "fmt"

func main() {
	// TODO: read integers in a loop with fmt.Scan until the user enters 0
	// Track the running `sum` and `count` of numbers entered before the 0.

	var n int
	sum := 0
	count := 0

	for {
		// TODO: use fmt.Scan to read the next integer into n
		// Stop the loop if the user entered 0 (or input ended).

		if n == 0 {
			break
		}
		sum += n
		count++
	}

	// TODO: print "Average: <value>" with the average as an integer.
}