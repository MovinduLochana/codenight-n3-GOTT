package main

import "testing"

func TestMean(t *testing.T) {
	tests := []struct {
		name   string
		scores []float64
		want   float64
	}{
		{"example", []float64{90, 86, 88}, 88},
		{"example2", []float64{10, 20}, 15},
		{"empty", []float64{}, 0},
		{"single", []float64{5}, 5},
		{"fractions", []float64{1, 2}, 1.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Mean(tt.scores...); got != tt.want {
				t.Errorf("Mean(%v) = %v; want %v", tt.scores, got, tt.want)
			}
		})
	}

	values := []float64{2, 4}
	if got := Mean(values...); got != 3 {
		t.Errorf("Mean(spread...) = %v; want 3", got)
	}
}