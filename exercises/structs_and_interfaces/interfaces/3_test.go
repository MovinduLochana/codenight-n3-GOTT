package main

import "testing"

func TestBiggest(t *testing.T) {
	shapes := []Shape{Rectangle{Width: 2, Height: 3}, Circle{Radius: 2}}
	got := Biggest(shapes)
	if got < 12.5 || got > 12.6 {
		t.Errorf("Biggest = %f; want ~12.56", got)
	}
}
