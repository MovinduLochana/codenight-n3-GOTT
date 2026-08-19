package main

import "testing"

func TestConcurrentSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"example", []int{1, 2, 3, 4}, 10},
		{"empty", []int{}, 0},
		{"single", []int{7}, 7},
		{"with negatives", []int{-1, 2, -3}, -2},
		{"larger", []int{1, 2, 3, 4, 5, 6}, 21},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < 5; i++ {
				if got := ConcurrentSum(tt.nums); got != tt.want {
					t.Fatalf("ConcurrentSum(%v) run %d = %d; want %d", tt.nums, i+1, got, tt.want)
				}
			}
		})
	}
}