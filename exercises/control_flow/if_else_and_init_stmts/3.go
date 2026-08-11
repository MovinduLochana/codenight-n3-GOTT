package main

import "fmt"

func ClassFare(age int, isStudent bool) int {
	// TODO: return the fare based on age and student status
	//   under 12 or 65+ → 25
	//   student         → 30 (unless already 25)
	//   everyone else   → 50
	return 0
}

func main() {
	fmt.Println(ClassFare(8, false))
	fmt.Println(ClassFare(70, false))
	fmt.Println(ClassFare(20, true))
	fmt.Println(ClassFare(20, false))
	fmt.Println(ClassFare(10, true))
}