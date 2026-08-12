package main

import "testing"

func TestIsPrime(t *testing.T) {
	tests := []struct {
		n        int
		expected bool
	}{
		{2, true},
		{7, true},
		{8, false},
		{1, false},
		{0, false},
		{13, true},
		{25, false},
	}
	for _, tt := range tests {
		got := IsPrime(tt.n)
		if got != tt.expected {
			t.Errorf("IsPrime(%d) = %t; want %t", tt.n, got, tt.expected)
		}
	}
}
