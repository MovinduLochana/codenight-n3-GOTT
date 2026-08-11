package main

import "fmt"

func IsLeapYear(year int) bool {
	// TODO: return true if year is a leap year
	// divisible by 4, except centuries unless divisible by 400
	return false
}

func main() {
	fmt.Println(IsLeapYear(2024)) // should print true
	fmt.Println(IsLeapYear(1900)) // should print false
	fmt.Println(IsLeapYear(2000)) // should print true
}
