package main

import (
	"reflect"
	"testing"
)

func TestTail(t *testing.T) {
	tests := []struct {
		s        []int
		n        int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{4, 5}},
		{[]int{1, 2}, 5, []int{1, 2}},
		{[]int{1, 2, 3}, 0, []int{}},
	}
	for _, tt := range tests {
		got := Tail(tt.s, tt.n)
		if got == nil && len(tt.expected) == 0 {
			continue
		}
		if !reflect.DeepEqual(got, tt.expected) {
			t.Errorf("Tail(%v, %d) = %v; want %v", tt.s, tt.n, got, tt.expected)
		}
	}
}
