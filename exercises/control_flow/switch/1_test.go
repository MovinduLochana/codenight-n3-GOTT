package main

import "testing"

func TestDayType(t *testing.T) {
	tests := []struct {
		day  string
		want string
	}{
		{"Sat", "Weekend"},
		{"Sun", "Weekend"},
		{"Mon", "Weekday"},
		{"Tue", "Weekday"},
		{"Wed", "Weekday"},
		{"Thu", "Weekday"},
		{"Fri", "Weekday"},
		{"saturday", "Weekday"},
		{"", "Weekday"},
	}
	for _, tt := range tests {
		if got := DayType(tt.day); got != tt.want {
			t.Errorf("DayType(%q) = %q; want %q", tt.day, got, tt.want)
		}
	}
}