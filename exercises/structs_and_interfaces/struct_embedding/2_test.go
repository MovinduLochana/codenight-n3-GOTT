package main

import "testing"

func TestWeightedRect(t *testing.T) {
	wr := WeightedRect{Rectangle: Rectangle{Width: 3, Height: 4}, Weight: 2}
	if got := wr.Area(); got != 24.0 {
		t.Errorf("Area = %f; want 24.0", got)
	}
}
