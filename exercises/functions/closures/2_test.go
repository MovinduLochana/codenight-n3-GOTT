package main

import "testing"

func TestMakeAdder(t *testing.T) {
	plus5 := MakeAdder(5)
	tests := []struct {
		n, want int
	}{
		{10, 15},
		{0, 5},
		{-5, 0},
		{100, 105},
	}
	for _, tt := range tests {
		if got := plus5(tt.n); got != tt.want {
			t.Errorf("plus5(%d) = %d; want %d", tt.n, got, tt.want)
		}
	}

	plus1 := MakeAdder(1)
	if got := plus1(1); got != 2 {
		t.Errorf("plus1(1) = %d; want 2 (closures must capture their own add)", got)
	}
}