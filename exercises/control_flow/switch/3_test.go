package main

import "testing"

func TestSeason(t *testing.T) {
	tests := []struct {
		m    int
		want string
	}{
		{12, "Winter"},
		{1, "Winter"},
		{2, "Winter"},
		{3, "Spring"},
		{4, "Spring"},
		{5, "Spring"},
		{6, "Summer"},
		{7, "Summer"},
		{8, "Summer"},
		{9, "Fall"},
		{10, "Fall"},
		{11, "Fall"},
		{0, "invalid"},
		{13, "invalid"},
		{-2, "invalid"},
	}
	for _, tt := range tests {
		if got := Season(tt.m); got != tt.want {
			t.Errorf("Season(%d) = %q; want %q", tt.m, got, tt.want)
		}
	}
}