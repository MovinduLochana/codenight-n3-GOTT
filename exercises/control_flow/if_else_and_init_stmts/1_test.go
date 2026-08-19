package main

import "testing"

func TestGrade(t *testing.T) {
	tests := []struct {
		score int
		want  string
	}{
		{95, "A"},
		{100, "A"},
		{90, "A"},
		{89, "B"},
		{82, "B"},
		{80, "B"},
		{79, "C"},
		{71, "C"},
		{70, "C"},
		{69, "F"},
		{50, "F"},
		{0, "F"},
	}
	for _, tt := range tests {
		if got := Grade(tt.score); got != tt.want {
			t.Errorf("Grade(%d) = %q; want %q", tt.score, got, tt.want)
		}
	}
}