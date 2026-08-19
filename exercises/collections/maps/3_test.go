package main

import (
	"reflect"
	"sort"
	"testing"
)

func sortKeys(in map[int][]string) map[int][]string {
	out := make(map[int][]string, len(in))
	for k, v := range in {
		s := append([]string(nil), v...)
		sort.Strings(s)
		out[k] = s
	}
	return out
}

func TestInvert(t *testing.T) {
	tests := []struct {
		name string
		m    map[string]int
		want map[int][]string
	}{
		{
			"example",
			map[string]int{"a": 1, "b": 1, "c": 2},
			map[int][]string{1: {"a", "b"}, 2: {"c"}},
		},
		{
			"all distinct",
			map[string]int{"x": 10, "y": 20},
			map[int][]string{10: {"x"}, 20: {"y"}},
		},
		{
			"empty",
			map[string]int{},
			map[int][]string{},
		},
		{
			"many to one",
			map[string]int{"p": 5, "q": 5, "r": 5},
			map[int][]string{5: {"p", "q", "r"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortKeys(Invert(tt.m))
			want := sortKeys(tt.want)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("Invert(%v) = %v; want %v", tt.m, got, tt.want)
			}
		})
	}
}