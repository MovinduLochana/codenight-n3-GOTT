package main

import "testing"

func TestArea(t *testing.T) {
	tests := []struct {
		name string
		r    Rectangle
		want float64
	}{
		{"example", Rectangle{Width: 3, Height: 4}, 12},
		{"square", Rectangle{Width: 5, Height: 5}, 25},
		{"zero", Rectangle{Width: 0, Height: 4}, 0},
		{"fractional", Rectangle{Width: 2.5, Height: 2}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Area(tt.r); got != tt.want {
				t.Errorf("Area(%v) = %v; want %v", tt.r, got, tt.want)
			}
		})
	}
}