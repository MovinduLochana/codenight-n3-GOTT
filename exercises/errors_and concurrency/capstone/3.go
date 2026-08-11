package main

import (
	"fmt"
	"sync"
)

func Validate(item string, qty int) error {
	if qty < 1 {
		return fmt.Errorf("invalid quantity %d for %s", qty, item)
	}
	return nil
}

func Checkout(items []string, qty []int) ([]int, error) {
	// TODO: validate every (item, qty) pair with Validate, return nil, err
	// if any is invalid.

	// TODO: launch one goroutine per item that computes qty[i] * 2 and
	// sends the result into a channel. Use a WaitGroup. Wait, then collect
	// all values with `for v := range ch`, SORT the result, and return it.
	return nil, nil
}

func main() {
	fmt.Println(Checkout([]string{"Coffee", "Tea"}, []int{2, 3}))
	fmt.Println(Checkout([]string{"Coffee"}, []int{0}))
}
