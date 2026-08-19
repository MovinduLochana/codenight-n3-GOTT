package main

import "testing"

func TestDaysInMonth(t *testing.T) {
	tests := []struct {
		m    int
		want int
	}{
		{2, 28},
		{4, 30},
		{1, 31},
		{13, 0},
		{0, 0},
		{7, 31},
		{6, 30},
		{11, 30},
		{12, 31},
		{3, 31},
		{9, 30},
		{8, 31},
		{5, 31},
		{10, 31},
		{14, 0},
		{-1, 0},
	}
	for _, tt := range tests {
		if got := DaysInMonth(tt.m); got != tt.want {
			t.Errorf("DaysInMonth(%d) = %d; want %d", tt.m, got, tt.want)
		}
	}
}