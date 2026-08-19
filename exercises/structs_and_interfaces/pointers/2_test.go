package main

import "testing"

func TestDouble(t *testing.T) {
	x := 7
	if got := Double(&x); got != 14 {
		t.Errorf("Double returned %d; want 14", got)
	}
	if x != 14 {
		t.Errorf("after Double, x = %d; want 14 (value must be written back)", x)
	}

	y := -3
	if got := Double(&y); got != -6 {
		t.Errorf("Double returned %d; want -6", got)
	}
	if y != -6 {
		t.Errorf("after Double, y = %d; want -6", y)
	}
}