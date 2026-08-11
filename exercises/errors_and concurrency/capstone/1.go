package main

import (
	"errors"
	"fmt"
)

func Validate(item string, qty int) error {
	// TODO: return an error when qty < 1, otherwise nil
	return nil
}

func main() {
	fmt.Println(Validate("Coffee", 2))
	fmt.Println(Validate("Tea", 0))
	fmt.Println(Validate("Tea", -3))
}
