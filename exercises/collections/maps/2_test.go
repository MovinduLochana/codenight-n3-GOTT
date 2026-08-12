package main

import (
	"reflect"
	"testing"
)

func TestMergeCounts(t *testing.T) {
	tests := []struct {
		name     string
		a        map[string]int
		b        map[string]int
		expected map[string]int
	}{
		{
			name:     "overlapping keys",
			a:        map[string]int{"go": 2, "is": 1},
			b:        map[string]int{"go": 1, "fun": 3},
			expected: map[string]int{"go": 3, "is": 1, "fun": 3},
		},
		{
			name:     "empty maps",
			a:        map[string]int{},
			b:        map[string]int{},
			expected: map[string]int{},
		},
		{
			name:     "disjoint keys",
			a:        map[string]int{"a": 1},
			b:        map[string]int{"b": 2},
			expected: map[string]int{"a": 1, "b": 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MergeCounts(tt.a, tt.b)
			if got == nil && len(tt.expected) == 0 {
				return
			}
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("MergeCounts(%v, %v) = %v; want %v", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}
