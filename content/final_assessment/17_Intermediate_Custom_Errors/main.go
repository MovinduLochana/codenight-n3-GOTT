package main

import (
	"errors"
	"fmt"
)

// TODO: Define ErrInvalidAge struct that holds an Age field (int) and implements the error interface.

// TODO: Implement VerifyAge(age int) error.

func main() {
	err := VerifyAge(-5)
	if err != nil {
		var invalidAgeErr *ErrInvalidAge
		if errors.As(err, &invalidAgeErr) {
			fmt.Printf("Error: age %d is invalid\n", invalidAgeErr.Age)
		} else {
			fmt.Println("Error:", err)
		}
	}
}
