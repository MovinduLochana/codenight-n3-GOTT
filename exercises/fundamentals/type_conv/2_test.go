package main

import "testing"

func TestCelsius(t *testing.T) {
	tests := []struct {
		f    float64
		want float64
	}{
		{68, 20},
		{32, 0},
		{212, 100},
		{-40, -40},
		{0, -17.77777777777778},
	}
	for _, tt := range tests {
		if got := Celsius(tt.f); got != tt.want {
			t.Errorf("Celsius(%v) = %v; want %v", tt.f, got, tt.want)
		}
	}
}