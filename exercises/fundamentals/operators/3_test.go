package main

import "testing"

func TestIsBetween(t *testing.T) {
	tests := []struct {
		n, low, high int
		want         bool
	}{
		{5, 1, 10, true},
		{0, 1, 10, false},
		{10, 1, 10, true},
		{1, 1, 1, true},
		{11, 1, 10, false},
		{-5, -10, 0, true},
	}
	for _, tt := range tests {
		if got := IsBetween(tt.n, tt.low, tt.high); got != tt.want {
			t.Errorf("IsBetween(%d, %d, %d) = %t; want %t", tt.n, tt.low, tt.high, got, tt.want)
		}
	}
}