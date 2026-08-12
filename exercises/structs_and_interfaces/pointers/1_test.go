package main

import "testing"

func TestSetFive(t *testing.T) {
	x := 0
	SetFive(&x)
	if x != 5 {
		t.Errorf("SetFive failed: x = %d; want 5", x)
	}
}
