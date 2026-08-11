package main

import (
	"errors"
	"fmt"
	"strconv"
)

func ParsePositive(s string) (int, error) {
	// TODO: parse s with strconv.Atoi.
	// Return the parse error unwrapped if it fails.
	// Return errors.New("number must be positive") if n < 0.
	// Otherwise return n, nil.
	return 0, nil
}

func main() {
	n, err := ParsePositive("42")
	fmt.Println(n, err)

	n, err = ParsePositive("-5")
	fmt.Println(n, err)

	n, err = ParsePositive("abc")
	fmt.Println(n, err)
}
