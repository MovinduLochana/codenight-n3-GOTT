package main

import "testing"

func TestIsLeapYear(t *testing.T) {
	tests := []struct {
		year int
		want bool
	}{
		{2024, true},
		{1900, false},
		{2000, true},
		{2100, false},
		{2023, false},
		{2004, true},
		{4, true},
		{100, false},
		{400, true},
	}
	for _, tt := range tests {
		if got := IsLeapYear(tt.year); got != tt.want {
			t.Errorf("IsLeapYear(%d) = %t; want %t", tt.year, got, tt.want)
		}
	}
}