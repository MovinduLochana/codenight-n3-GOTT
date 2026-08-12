package main

import "testing"

func TestFactorial(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{5, 120},
		{3, 6},
		{0, 1},
		{1, 1},
	}
	for _, tt := range tests {
		got := Factorial(tt.n)
		if got != tt.expected {
			t.Errorf("Factorial(%d) = %d; want %d", tt.n, got, tt.expected)
		}
	}
}
