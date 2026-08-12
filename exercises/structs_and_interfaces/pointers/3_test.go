package main

import "testing"

func TestSwapNums(t *testing.T) {
	a, b := 1, 2
	SwapNums(&a, &b)
	if a != 2 || b != 1 {
		t.Errorf("SwapNums failed: a=%d, b=%d", a, b)
	}
}
