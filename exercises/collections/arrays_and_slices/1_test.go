package main

import (
	"reflect"
	"testing"
)

func TestAppendThree(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{}, []int{1, 2, 3}},
		{[]int{9}, []int{9, 1, 2, 3}},
	}
	for _, tt := range tests {
		got := AppendThree(tt.input)
		if !reflect.DeepEqual(got, tt.expected) {
			t.Errorf("AppendThree(%v) = %v; want %v", tt.input, got, tt.expected)
		}
	}
}
