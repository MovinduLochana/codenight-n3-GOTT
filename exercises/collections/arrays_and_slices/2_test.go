package main

import (
	"reflect"
	"testing"
)

func TestRange(t *testing.T) {
	tests := []struct {
		n        int
		expected []int
	}{
		{5, []int{0, 1, 2, 3, 4}},
		{0, []int{}},
	}
	for _, tt := range tests {
		got := Range(tt.n)
		if got == nil && len(tt.expected) == 0 {
			continue
		}
		if !reflect.DeepEqual(got, tt.expected) {
			t.Errorf("Range(%d) = %v; want %v", tt.n, got, tt.expected)
		}
	}
}
