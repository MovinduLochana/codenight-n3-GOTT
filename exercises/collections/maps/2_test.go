package main

import (
	"reflect"
	"testing"
)

func TestMergeCounts(t *testing.T) {
	tests := []struct {
		name string
		a, b map[string]int
		want map[string]int
	}{
		{
			"example",
			map[string]int{"go": 2, "is": 1},
			map[string]int{"go": 1, "fun": 3},
			map[string]int{"go": 3, "is": 1, "fun": 3},
		},
		{
			"disjoint",
			map[string]int{"a": 1},
			map[string]int{"b": 2},
			map[string]int{"a": 1, "b": 2},
		},
		{
			"empty a",
			map[string]int{},
			map[string]int{"x": 5},
			map[string]int{"x": 5},
		},
		{
			"empty b",
			map[string]int{"x": 5},
			map[string]int{},
			map[string]int{"x": 5},
		},
		{
			"both empty",
			map[string]int{},
			map[string]int{},
			map[string]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeCounts(tt.a, tt.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeCounts(%v, %v) = %v; want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}