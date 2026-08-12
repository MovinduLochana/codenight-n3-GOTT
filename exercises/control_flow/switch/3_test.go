package main

import "testing"

func TestSeason(t *testing.T) {
	tests := []struct {
		m    int
		want string
	}{
		{1, "Winter"},
		{4, "Spring"},
		{7, "Summer"},
		{10, "Fall"},
		{13, "invalid"},
	}
	for _, tt := range tests {
		if got := Season(tt.m); got != tt.want {
			t.Errorf("Season(%d) = %q; want %q", tt.m, got, tt.want)
		}
	}
}
