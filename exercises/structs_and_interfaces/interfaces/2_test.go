package main

import "testing"

func TestTotalPerimeter(t *testing.T) {
	shapes := []Perimeterer{Rectangle{Width: 3, Height: 4}}
	if got := TotalPerimeter(shapes); got != 14.0 {
		t.Errorf("TotalPerimeter = %f; want 14.0", got)
	}
}
