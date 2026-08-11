package main

import "fmt"

func Double(ptr *int) int {
	// TODO: double the value pointed to by ptr, write it back, and return it
	return 0
}

func main() {
	x := 7
	fmt.Println(Double(&x)) // 14
	fmt.Println(x)          // 14
}