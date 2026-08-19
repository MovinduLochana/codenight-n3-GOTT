package main

import (
	"math"
	"testing"
)

func TestTotalPerimeter(t *testing.T) {
	shapes := []Perimeterer{
		Rectangle{Width: 2, Height: 3},
		Circle{Radius: 1},
	}
	want := 2*(2+3) + 2*3.14159*1
	if got := TotalPerimeter(shapes); math.Abs(got-want) > 1e-9 {
		t.Errorf("TotalPerimeter(shapes) = %v; want %v", got, want)
	}

	if got := TotalPerimeter([]Perimeterer{}); got != 0 {
		t.Errorf("TotalPerimeter(empty) = %v; want 0", got)
	}

	shapes2 := []Perimeterer{
		Rectangle{Width: 5, Height: 5},
	}
	if got := TotalPerimeter(shapes2); got != 20 {
		t.Errorf("TotalPerimeter(square rect) = %v; want 20", got)
	}
}