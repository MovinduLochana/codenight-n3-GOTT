package main

import "testing"

func TestTotalAreaInterface(t *testing.T) {
	shapes := []Shape{Rectangle{Width: 3, Height: 4}, Circle{Radius: 1}}
	got := TotalArea(shapes)
	if got < 15.1 || got > 15.2 {
		t.Errorf("TotalArea = %f; want ~15.14", got)
	}
}
