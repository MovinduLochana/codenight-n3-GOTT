package main

import "testing"

func sameStrIntMap(a, b map[string]int) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}

func TestWordCount(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		want  map[string]int
	}{
		{
			"example",
			[]string{"go", "is", "fun", "go", "is", "go"},
			map[string]int{"go": 3, "is": 2, "fun": 1},
		},
		{
			"empty",
			[]string{},
			map[string]int{},
		},
		{
			"repeats",
			[]string{"a", "a", "a"},
			map[string]int{"a": 3},
		},
		{
			"single",
			[]string{"hello"},
			map[string]int{"hello": 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WordCount(tt.words); !sameStrIntMap(got, tt.want) {
				t.Errorf("WordCount(%v) = %v; want %v", tt.words, got, tt.want)
			}
		})
	}
}