package main

import "testing"

func TestScale(t *testing.T) {
	r := Rectangle{Width: 3, Height: 4}
	r.Scale(2)
	if r.Width != 6 || r.Height != 8 {
		t.Errorf("after Scale(2), rect = %v; want Width 6, Height 8", r)
	}

	r2 := Rectangle{Width: 10, Height: 1}
	r2.Scale(0.5)
	if r2.Width != 5 || r2.Height != 0.5 {
		t.Errorf("after Scale(0.5), rect = %v; want Width 5, Height 0.5", r2)
	}
}