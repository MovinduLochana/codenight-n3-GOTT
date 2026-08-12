package main

import "testing"

func TestDouble(t *testing.T) {
	x := 7
	got := Double(&x)
	if got != 14 || x != 14 {
		t.Errorf("Double failed: got %d, x = %d", got, x)
	}
}
