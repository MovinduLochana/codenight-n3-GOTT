package main

import "testing"

func TestToPixels(t *testing.T) {
	tests := []struct {
		inches, dpi float64
		want        int
	}{
		{1.5, 300, 450},
		{0.5, 72, 36},
		{2, 96, 192},
		{1, 72, 72},
		{0, 100, 0},
	}
	for _, tt := range tests {
		if got := ToPixels(tt.inches, tt.dpi); got != tt.want {
			t.Errorf("ToPixels(%v, %v) = %d; want %d", tt.inches, tt.dpi, got, tt.want)
		}
	}
}