package main

import "testing"

func TestArea(t *testing.T) {
	r := Rectangle{Width: 3, Height: 4}
	if got := Area(r); got != 12.0 {
		t.Errorf("Area = %f; want 12.0", got)
	}
}
