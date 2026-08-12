package main

import "testing"

func TestScale(t *testing.T) {
	r := Rectangle{Width: 3, Height: 4}
	r.Scale(2)
	if r.Width != 6 || r.Height != 8 {
		t.Errorf("Scale failed: %v", r)
	}
}
