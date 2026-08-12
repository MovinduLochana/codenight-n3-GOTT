package main

import (
	"errors"
	"fmt"
)

// TODO: Implement the Divide function here.
// It should accept two float64 parameters and return (float64, error).

func main() {
	res1, err1 := Divide(10.0, 2.0)
	if err1 == nil {
		fmt.Printf("Result: %.2f\n", res1)
	}

	_, err2 := Divide(5.0, 0.0)
	if err2 != nil {
		fmt.Printf("Error: %v\n", err2)
	}
}
