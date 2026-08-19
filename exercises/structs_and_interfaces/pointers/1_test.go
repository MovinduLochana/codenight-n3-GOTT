package main

import "testing"

func TestSetFive(t *testing.T) {
	x := 3
	SetFive(&x)
	if x != 5 {
		t.Errorf("after SetFive, x = %d; want 5", x)
	}

	y := 100
	SetFive(&y)
	if y != 5 {
		t.Errorf("after SetFive, y = %d; want 5", y)
	}
}