package main

import (
	"errors"
	"fmt"
	"strconv"
)

func ReadAge(s string) (int, error) {
	// TODO: parse s with strconv.Atoi.
	// On failure, wrap the error: fmt.Errorf("invalid age: %w", err)
	return 0, nil
}

func main() {
	n, err := ReadAge("25")
	fmt.Println(n, err)

	_, err = ReadAge("xx")
	fmt.Println(err)                               // invalid age: strconv.Atoi: ...
	fmt.Println(errors.Is(err, strconv.ErrSyntax)) // true
}
