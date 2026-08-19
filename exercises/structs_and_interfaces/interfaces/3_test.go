package main

import (
	"math"
	"testing"
)

func TestBiggest(t *testing.T) {
	shapes := []Shape{
		Rectangle{Width: 2, Height: 3},
		Circle{Radius: 2},
	}
	want := 3.14159 * 2 * 2
	if got := Biggest(shapes); math.Abs(got-want) > 1e-9 {
		t.Errorf("Biggest(shapes) = %v; want %v", got, want)
	}

	single := []Shape{Rectangle{Width: 4, Height: 2}}
	if got := Biggest(single); got != 8 {
		t.Errorf("Biggest(single) = %v; want 8", got)
	}

	if got := Biggest([]Shape{}); got != 0 {
		t.Errorf("Biggest(empty) = %v; want 0", got)
	}
}