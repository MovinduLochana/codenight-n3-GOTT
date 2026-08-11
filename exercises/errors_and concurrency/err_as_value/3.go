package main

import (
	"errors"
	"fmt"
	"strconv"
)

// TODO: declare `var ErrNotPositive = errors.New("number must be positive")`

func ParsePositive(s string) (int, error) {
	// TODO: parse s with strconv.Atoi (return parse error unwrapped).
	// If n <= 0, return 0, ErrNotPositive. Otherwise return n, nil.
	return 0, nil
}

func main() {
	n, err := ParsePositive("7")
	fmt.Println(n, err)
	fmt.Println(errors.Is(err, ErrNotPositive)) // false

	_, err = ParsePositive("-5")
	fmt.Println(errors.Is(err, ErrNotPositive)) // true

	_, err = ParsePositive("0")
	fmt.Println(errors.Is(err, ErrNotPositive)) // true

	_, err = ParsePositive("abc")
	fmt.Println(err)                            // strconv.Atoi error
	fmt.Println(errors.Is(err, ErrNotPositive)) // false
}
