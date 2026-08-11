package main

import "fmt"

func Season(m int) string {
	// TODO: use an expression-less switch with grouped cases (e.g. case 12, 1, 2:)
	return ""
}

func main() {
	fmt.Println(Season(12)) // Winter
	fmt.Println(Season(4))  // Spring
	fmt.Println(Season(8))  // Summer
	fmt.Println(Season(10)) // Fall
	fmt.Println(Season(0))  // invalid
}