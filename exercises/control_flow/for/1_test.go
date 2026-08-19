package main

import "testing"

func TestSumEvens(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{10, 30}, // 2+4+6+8+10
		{5, 6},   // 2+4
		{0, 0},
		{1, 0},
		{2, 2},
		{3, 2},
		{6, 12},
		{20, 110},
	}
	for _, tt := range tests {
		if got := SumEvens(tt.n); got != tt.want {
			t.Errorf("SumEvens(%d) = %d; want %d", tt.n, got, tt.want)
		}
	}
}