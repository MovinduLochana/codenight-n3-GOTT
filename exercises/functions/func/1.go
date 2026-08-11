package main

import (
	"errors"
	"fmt"
)

func Divide(a, b int) (int, error) {
	// TODO: return an error if b == 0, otherwise return a / b and nil
	return 0, nil
}

func main() {
	result, err := Divide(10, 2)
	fmt.Println(result, err)

	result, err = Divide(5, 0)
	fmt.Println(result, err)
}