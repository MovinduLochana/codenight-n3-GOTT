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

func TestIndexAll(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{"example", []int{1, 2, 1, 4, 1}, 1, []int{0, 2, 4}},
		{"absent", []int{1, 2, 3}, 9, []int{}},
		{"empty", []int{}, 0, []int{}},
		{"single", []int{5}, 5, []int{0}},
		{"all", []int{2, 2, 2}, 2, []int{0, 1, 2}},
		{"adjacent", []int{1, 1, 2, 1}, 1, []int{0, 1, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IndexAll(tt.nums, tt.target); !sameInts(got, tt.want) {
				t.Errorf("IndexAll(%v, %d) = %v; want %v", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}