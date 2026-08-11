package main

import "fmt"

func MergeCounts(a, b map[string]int) map[string]int {
	// TODO: combine a and b into a new map, summing counts
	return nil
}

func main() {
	fmt.Println(MergeCounts(
		map[string]int{"go": 2, "is": 1},
		map[string]int{"go": 1, "fun": 3},
	))
}