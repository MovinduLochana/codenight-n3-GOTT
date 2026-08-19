package main

import (
	"math"
	"testing"
)

func TestTotalAreaEmbedding(t *testing.T) {
	sq := Square{Rectangle: Rectangle{Width: 3, Height: 4}, Side: 5}
	if got := sq.Area(); got != 25 {
		t.Errorf("Square.Area() = %v; want 25 (must shadow Rectangle.Area)", got)
	}
	if got := sq.Rectangle.Area(); got != 12 {
		t.Errorf("Rectangle.Area() = %v; want 12", got)
	}

	shapes := []Shape{
		Square{Rectangle: Rectangle{Width: 3, Height: 4}, Side: 5},
		Rectangle{Width: 2, Height: 3},
	}
	want := 25.0 + 6.0
	if got := TotalArea(shapes); math.Abs(got-want) > 1e-9 {
		t.Errorf("TotalArea(shapes) = %v; want %v", got, want)
	}

	if got := TotalArea([]Shape{}); got != 0 {
		t.Errorf("TotalArea(empty) = %v; want 0", got)
	}
}