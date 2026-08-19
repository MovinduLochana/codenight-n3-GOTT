package main

import "testing"

func TestSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"example", []int{1, 2, 3, 4}, 10},
		{"empty", []int{}, 0},
		{"single", []int{5}, 5},
		{"with negatives", []int{-1, -2, 3}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.nums); got != tt.want {
				t.Errorf("Sum(%v) = %d; want %d", tt.nums, got, tt.want)
			}
		})
	}
}