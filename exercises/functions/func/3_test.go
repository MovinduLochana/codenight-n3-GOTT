package main

import "testing"

func TestMinMax(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		wantMin  int
		wantMax  int
	}{
		{"example", []int{3, 1, 4, 1, 5}, 1, 5},
		{"single", []int{7}, 7, 7},
		{"with negatives", []int{-2, -5, 0}, -5, 0},
		{"reversed", []int{9, 8, 7, 6}, 6, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lo, hi := MinMax(tt.nums)
			if lo != tt.wantMin || hi != tt.wantMax {
				t.Errorf("MinMax(%v) = (%d, %d); want (%d, %d)", tt.nums, lo, hi, tt.wantMin, tt.wantMax)
			}
		})
	}
}