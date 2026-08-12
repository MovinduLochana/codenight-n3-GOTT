package main

import (
	"reflect"
	"testing"
)

func TestWordCount(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected map[string]int
	}{
		{
			name:  "standard words",
			input: []string{"go", "is", "fun", "go", "is", "go"},
			expected: map[string]int{
				"go":  3,
				"is":  2,
				"fun": 1,
			},
		},
		{
			name:     "empty slice",
			input:    []string{},
			expected: map[string]int{},
		},
		{
			name:  "single word repeated",
			input: []string{"gopher", "gopher", "gopher"},
			expected: map[string]int{
				"gopher": 3,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := WordCount(tt.input)
			if got == nil && len(tt.expected) == 0 {
				return
			}
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("WordCount(%v) = %v; want %v", tt.input, got, tt.expected)
			}
		})
	}
}
