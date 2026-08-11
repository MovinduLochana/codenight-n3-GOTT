package main

import "fmt"

func Invert(m map[string]int) map[int][]string {
	// TODO: return a map whose keys are m's values and whose values
	// are slices of m's keys that mapped to that value
	return nil
}

func main() {
	fmt.Println(Invert(map[string]int{"a": 1, "b": 1, "c": 2}))
}