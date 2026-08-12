package main

import "testing"

func TestSumEvens(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{10, 30},
		{5, 6},
		{0, 0},
		{1, 0},
		{2, 2},
	}
	for _, tt := range tests {
		got := SumEvens(tt.n)
		if got != tt.expected {
			t.Errorf("SumEvens(%d) = %d; want %d", tt.n, got, tt.expected)
		}
	}
}
