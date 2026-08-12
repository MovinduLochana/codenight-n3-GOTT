package main

import "testing"

func TestNamedRectangle(t *testing.T) {
	nr := NamedRectangle{Rectangle: Rectangle{Width: 3, Height: 4}, Name: "Box"}
	if got := nr.Area(); got != 12.0 {
		t.Errorf("Area = %f; want 12.0", got)
	}
}
