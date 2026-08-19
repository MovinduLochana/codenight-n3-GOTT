package main

import "testing"

func TestWeightedRect(t *testing.T) {
	wr := WeightedRect{Rectangle: Rectangle{Width: 3, Height: 4}, Weight: 2}
	if got := wr.Area(); got != 24 {
		t.Errorf("WeightedRect.Area() = %v; want 24 (area * weight)", got)
	}
	if got := wr.Rectangle.Area(); got != 12 {
		t.Errorf("Rectangle.Area() = %v; want 12 (promoted method must still work)", got)
	}
}