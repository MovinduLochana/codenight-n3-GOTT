package main

import "testing"

func TestJoin(t *testing.T) {
	tests := []struct {
		name  string
		sep   string
		parts []string
		want  string
	}{
		{"example", ", ", []string{"go", "is", "fun"}, "go, is, fun"},
		{"single part", "|", []string{"a"}, "a"},
		{"no parts", "-", []string{}, ""},
		{"empty sep", "", []string{"a", "b"}, "ab"},
		{"pipes", "|", []string{"a", "b", "c"}, "a|b|c"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Join(tt.sep, tt.parts...); got != tt.want {
				t.Errorf("Join(%q, %v) = %q; want %q", tt.sep, tt.parts, got, tt.want)
			}
		})
	}
}