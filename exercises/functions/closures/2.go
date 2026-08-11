package main

import "fmt"

func MakeAdder(add int) func(int) int {
	// TODO: return a closure that captures `add` and adds it to its argument
	return nil
}

func main() {
	plus5 := MakeAdder(5)
	fmt.Println(plus5(10)) // 15
	fmt.Println(plus5(0))  // 5
}