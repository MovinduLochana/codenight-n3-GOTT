package main

import "testing"

func TestAverage(t *testing.T) {
	tests := []struct {
		a, b int
		want float64
	}{
		{3, 4, 3.5},
		{0, 0, 0},
		{5, 2, 3.5},
		{1, 2, 1.5},
		{10, 5, 7.5},
	}
	for _, tt := range tests {
		if got := Average(tt.a, tt.b); got != tt.want {
			t.Errorf("Average(%d, %d) = %f; want %f", tt.a, tt.b, got, tt.want)
		}
	}
}