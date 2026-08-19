package main

import (
	"math"
	"testing"
)

func TestTotalAreaInterface(t *testing.T) {
	c := Circle{Radius: 1}
	if got := c.Area(); math.Abs(got-3.14159) > 1e-9 {
		t.Errorf("Circle{1}.Area() = %v; want 3.14159", got)
	}

	shapes := []Shape{
		Rectangle{Width: 2, Height: 3},
		Circle{Radius: 1},
	}
	want := 6.0 + 3.14159
	if got := TotalArea(shapes); math.Abs(got-want) > 1e-9 {
		t.Errorf("TotalArea(shapes) = %v; want %v", got, want)
	}

	if got := TotalArea([]Shape{}); got != 0 {
		t.Errorf("TotalArea(empty) = %v; want 0", got)
	}
}