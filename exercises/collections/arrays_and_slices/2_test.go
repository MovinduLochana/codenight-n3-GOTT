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

func TestRange(t *testing.T) {
	tests := []struct {
		n    int
		want []int
	}{
		{5, []int{0, 1, 2, 3, 4}},
		{0, []int{}},
		{1, []int{0}},
		{3, []int{0, 1, 2}},
		{6, []int{0, 1, 2, 3, 4, 5}},
	}
	for _, tt := range tests {
		if got := Range(tt.n); !sameInts(got, tt.want) {
			t.Errorf("Range(%d) = %v; want %v", tt.n, got, tt.want)
		}
	}
}