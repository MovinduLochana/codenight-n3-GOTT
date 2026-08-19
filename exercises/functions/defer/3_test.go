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

func TestLoopCleanup(t *testing.T) {
	tests := []struct {
		n    int
		want []int
	}{
		{3, []int{2, 1, 0}},
		{1, []int{0}},
		{0, []int{}},
		{5, []int{4, 3, 2, 1, 0}},
	}
	for _, tt := range tests {
		if got := LoopCleanup(tt.n); !sameInts(got, tt.want) {
			t.Errorf("LoopCleanup(%d) = %v; want %v", tt.n, got, tt.want)
		}
	}
}