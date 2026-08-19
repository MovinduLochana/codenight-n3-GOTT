package main

import "testing"

func TestIsEven(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{4, true},
		{7, false},
		{0, true},
		{-2, true},
		{-1, false},
		{1, false},
		{100, true},
	}
	for _, tt := range tests {
		if got := IsEven(tt.n); got != tt.want {
			t.Errorf("IsEven(%d) = %t; want %t", tt.n, got, tt.want)
		}
	}
}