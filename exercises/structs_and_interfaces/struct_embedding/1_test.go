package main

import "testing"

func TestNamedRectangle(t *testing.T) {
	nr := NamedRectangle{Rectangle: Rectangle{Width: 3, Height: 4}, Name: "MyRect"}
	if nr.Name != "MyRect" {
		t.Errorf("Name = %q; want %q", nr.Name, "MyRect")
	}
	if got := nr.Area(); got != 12 {
		t.Errorf("Area() = %v; want 12 (promoted Rectangle.Area)", got)
	}
	if nr.Width != 3 || nr.Height != 4 {
		t.Errorf("promoted fields not accessible: %v", nr)
	}
}