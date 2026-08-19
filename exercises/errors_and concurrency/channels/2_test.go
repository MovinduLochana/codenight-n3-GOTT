package main

import "testing"

func sameInts(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestSquaresPipeline(t *testing.T) {
	tests := []struct {
		k    int
		want []int
	}{
		{4, []int{1, 4, 9, 16}},
		{0, []int{}},
		{1, []int{1}},
		{3, []int{1, 4, 9}},
	}
	for _, tt := range tests {
		if got := SquaresPipeline(tt.k); !sameInts(got, tt.want) {
			t.Errorf("SquaresPipeline(%d) = %v; want %v", tt.k, got, tt.want)
		}
	}
}