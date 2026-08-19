package main

import "testing"

func TestFactorial(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{5, 120},
		{3, 6},
		{0, 1},
		{1, 1},
		{2, 2},
		{4, 24},
		{6, 720},
	}
	for _, tt := range tests {
		if got := Factorial(tt.n); got != tt.want {
			t.Errorf("Factorial(%d) = %d; want %d", tt.n, got, tt.want)
		}
	}
}