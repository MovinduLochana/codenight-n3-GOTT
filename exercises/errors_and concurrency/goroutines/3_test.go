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

func TestSquares(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"example", []int{4, 1, 3}, []int{1, 9, 16}},
		{"empty", []int{}, []int{}},
		{"single", []int{2}, []int{4}},
		{"zero", []int{0, 3}, []int{0, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < 5; i++ {
				if got := Squares(tt.nums); !sameInts(got, tt.want) {
					t.Fatalf("Squares(%v) run %d = %v; want %v", tt.nums, i+1, got, tt.want)
				}
			}
		})
	}
}