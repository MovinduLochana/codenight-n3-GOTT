package main

import "testing"

func TestIsPrime(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{2, true},
		{7, true},
		{8, false},
		{1, false},
		{0, false},
		{13, true},
		{25, false},
		{3, true},
		{9, false},
		{17, true},
		{-1, false},
		{2, true},
	}
	for _, tt := range tests {
		if got := IsPrime(tt.n); got != tt.want {
			t.Errorf("IsPrime(%d) = %t; want %t", tt.n, got, tt.want)
		}
	}
}