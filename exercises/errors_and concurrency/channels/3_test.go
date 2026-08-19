package main

import "testing"

func TestEvenSum(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{10, 30},
		{5, 6},
		{0, 0},
		{1, 0},
		{2, 2},
		{6, 12},
		{100, 2550},
	}
	for _, tt := range tests {
		if got := EvenSum(tt.n); got != tt.want {
			t.Errorf("EvenSum(%d) = %d; want %d", tt.n, got, tt.want)
		}
	}
}