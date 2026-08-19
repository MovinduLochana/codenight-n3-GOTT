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

func TestCut(t *testing.T) {
	tests := []struct {
		name    string
		s       []int
		lo, hi  int
		want    []int
	}{
		{"example", []int{10, 20, 30, 40, 50}, 1, 3, []int{20, 30}},
		{"example2", []int{5, 6, 7}, 0, 2, []int{5, 6}},
		{"empty window", []int{1, 2, 3}, 1, 1, []int{}},
		{"full slice", []int{1, 2, 3}, 0, 3, []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Cut(tt.s, tt.lo, tt.hi); !sameInts(got, tt.want) {
				t.Errorf("Cut(%v, %d, %d) = %v; want %v", tt.s, tt.lo, tt.hi, got, tt.want)
			}
		})
	}
}