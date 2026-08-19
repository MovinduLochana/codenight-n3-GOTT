package main

import "testing"

func TestContains(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   bool
	}{
		{"found", []int{1, 2, 3}, 2, true},
		{"not found", []int{1, 2, 3}, 9, false},
		{"empty", []int{}, 0, false},
		{"first element", []int{7, 8, 9}, 7, true},
		{"last element", []int{7, 8, 9}, 9, true},
		{"single not found", []int{5}, 6, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.nums, tt.target); got != tt.want {
				t.Errorf("Contains(%v, %d) = %t; want %t", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}