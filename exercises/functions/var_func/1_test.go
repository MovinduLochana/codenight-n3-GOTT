package main

import "testing"

func TestMax(t *testing.T) {
	if got := Max(3, 7, 2); got != 7 {
		t.Errorf("Max(3, 7, 2) = %d; want 7", got)
	}
	if got := Max(); got != 0 {
		t.Errorf("Max() = %d; want 0", got)
	}
}
