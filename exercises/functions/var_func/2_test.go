package main

import "testing"

func TestJoin(t *testing.T) {
	expected := "go, is, fun"
	if got := Join(", ", "go", "is", "fun"); got != expected {
		t.Errorf("Join = %q; want %q", got, expected)
	}
}
