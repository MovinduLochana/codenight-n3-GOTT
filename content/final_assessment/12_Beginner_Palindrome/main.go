package main

import (
	"fmt"
	"strings"
)

// TODO: Implement IsPalindrome(s string) bool.
// It should return true if `s` is a palindrome (reads same forwards and backwards), ignoring case.

func main() {
	fmt.Printf("Racecar: %t\n", IsPalindrome("Racecar"))
	fmt.Printf("golang: %t\n", IsPalindrome("golang"))
}
