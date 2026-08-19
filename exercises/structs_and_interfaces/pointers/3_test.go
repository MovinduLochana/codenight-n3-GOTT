package main

import "testing"

func TestSwapNums(t *testing.T) {
	x, y := 1, 2
	SwapNums(&x, &y)
	if x != 2 || y != 1 {
		t.Errorf("after SwapNums, x=%d y=%d; want x=2 y=1", x, y)
	}

	a, b := -5, 10
	SwapNums(&a, &b)
	if a != 10 || b != -5 {
		t.Errorf("after SwapNums, a=%d b=%d; want a=10 b=-5", a, b)
	}

	equalA, equalB := 3, 3
	SwapNums(&equalA, &equalB)
	if equalA != 3 || equalB != 3 {
		t.Errorf("swapping equal values changed them: %d %d", equalA, equalB)
	}
}