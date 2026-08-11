package main

import "fmt"

func SwapNums(a, b *int) {
	// TODO: exchange the values pointed to by a and b
	// use a temporary variable
}

func main() {
	x, y := 1, 2
	SwapNums(&x, &y)
	fmt.Println(x, y) // 2 1
}