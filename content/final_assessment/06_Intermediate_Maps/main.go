package main

import "fmt"

// TODO: Implement CharFrequency(s string) map[rune]int
// It should skip space characters ' '.

func main() {
	input := "golang programming"
	freq := CharFrequency(input)

	fmt.Printf("g: %d\n", freq['g'])
	fmt.Printf("o: %d\n", freq['o'])
	fmt.Printf("l: %d\n", freq['l'])
	fmt.Printf("space: %d\n", freq[' '])
}
