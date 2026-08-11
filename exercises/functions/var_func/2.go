package main

import "fmt"

func Join(sep string, parts ...string) string {
	// TODO: concatenate parts with sep between them (never at the ends)
	return ""
}

func main() {
	fmt.Println(Join(", ", "go", "is", "fun"))
	fmt.Println(Join("|", "a"))
	fmt.Println(Join("-"))
}