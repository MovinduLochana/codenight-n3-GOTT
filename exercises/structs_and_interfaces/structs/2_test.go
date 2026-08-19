package main

import "testing"

func TestVolume(t *testing.T) {
	tests := []struct {
		name string
		b    Box
		want float64
	}{
		{"example", Box{Length: 2, Width: 3, Height: 4}, 24},
		{"unit", Box{Length: 1, Width: 1, Height: 1}, 1},
		{"zero", Box{Length: 0, Width: 3, Height: 4}, 0},
		{"fractional", Box{Length: 1.5, Width: 2, Height: 2}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Volume(tt.b); got != tt.want {
				t.Errorf("Volume(%v) = %v; want %v", tt.b, got, tt.want)
			}
		})
	}
}