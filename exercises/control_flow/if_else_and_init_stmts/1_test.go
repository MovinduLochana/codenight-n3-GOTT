package main

import "testing"

func TestGrade(t *testing.T) {
	tests := []struct {
		score int
		want  string
	}{
		{95, "A"},
		{82, "B"},
		{71, "C"},
		{50, "F"},
	}
	for _, tt := range tests {
		if got := Grade(tt.score); got != tt.want {
			t.Errorf("Grade(%d) = %q; want %q", tt.score, got, tt.want)
		}
	}
}
