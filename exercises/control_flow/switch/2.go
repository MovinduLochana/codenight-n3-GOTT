package main

import "fmt"

func DaysInMonth(m int) int {
	// TODO: use a switch with grouped cases to return the day count
	return 0
}

func main() {
	fmt.Println(DaysInMonth(2))  // 28
	fmt.Println(DaysInMonth(4))  // 30
	fmt.Println(DaysInMonth(1))  // 31
	fmt.Println(DaysInMonth(13)) // 0
}