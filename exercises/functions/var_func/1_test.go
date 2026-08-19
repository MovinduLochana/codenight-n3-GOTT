package main

import "testing"

func TestMax(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"example", []int{3, 7, 2}, 7},
		{"single", []int{5}, 5},
		{"empty", []int{}, 0},
		{"negative", []int{-1, -5}, -1},
		{"duplicates", []int{4, 4, 3}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.nums...); got != tt.want {
				t.Errorf("Max(%v) = %d; want %d", tt.nums, got, tt.want)
			}
		})
	}

	spread := []int{2, 4}
	if got := Max(spread...); got != 4 {
		t.Errorf("Max(spread...) = %d; want 4", got)
	}
}