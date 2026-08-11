package main

import (
	"errors"
	"fmt"
)

func DivMod(a, b int) (quotient, remainder int, err error) {
	// TODO: return quotient = a/b, remainder = a%b, err = nil when b != 0;
	// return 0, 0, errors.New(...) when b == 0
	return 0, 0, nil
}

func main() {
	q, r, err := DivMod(10, 3)
	fmt.Println(q, r, err)

	q, r, err = DivMod(10, 0)
	fmt.Println(q, r, err)
}