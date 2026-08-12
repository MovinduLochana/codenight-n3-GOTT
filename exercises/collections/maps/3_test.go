package main

import (
	"sort"
	"testing"
)

func TestInvert(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		expected map[int][]string
	}{
		{
			name:  "multiple keys same value",
			input: map[string]int{"a": 1, "b": 1, "c": 2},
			expected: map[int][]string{
				1: {"a", "b"},
				2: {"c"},
			},
		},
		{
			name:     "empty map",
			input:    map[string]int{},
			expected: map[int][]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Invert(tt.input)
			if got == nil && len(tt.expected) == 0 {
				return
			}
			if len(got) != len(tt.expected) {
				t.Fatalf("Invert(%v) length = %d; want %d", tt.input, len(got), len(tt.expected))
			}
			for k, expectedSlice := range tt.expected {
				gotSlice, ok := got[k]
				if !ok {
					t.Errorf("Invert(%v) missing key %d", tt.input, k)
					continue
				}
				gotCopy := append([]string(nil), gotSlice...)
				expectedCopy := append([]string(nil), expectedSlice...)
				sort.Strings(gotCopy)
				sort.Strings(expectedCopy)
				if len(gotCopy) != len(expectedCopy) {
					t.Errorf("Invert(%v)[%d] = %v; want %v", tt.input, k, gotSlice, expectedSlice)
					continue
				}
				for i := range gotCopy {
					if gotCopy[i] != expectedCopy[i] {
						t.Errorf("Invert(%v)[%d] = %v; want %v", tt.input, k, gotSlice, expectedSlice)
						break
					}
				}
			}
		})
	}
}
