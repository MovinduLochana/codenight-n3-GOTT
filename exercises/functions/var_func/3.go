package main

import "fmt"

func Mean(scores ...float64) float64 {
	// TODO: return the average of all scores, or 0 if empty
	return 0
}

func main() {
	fmt.Println(Mean(90, 86, 88))
	fmt.Println(Mean(10, 20))
	fmt.Println(Mean())

	values := []float64{2, 4}
	fmt.Println(Mean(values...)) // spread a slice into the variadic call
}