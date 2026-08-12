package main

import "testing"

func TestClassFare(t *testing.T) {
	tests := []struct {
		age       int
		isStudent bool
		want      int
	}{
		{8, false, 25},
		{70, false, 25},
		{20, true, 30},
		{20, false, 50},
	}
	for _, tt := range tests {
		if got := ClassFare(tt.age, tt.isStudent); got != tt.want {
			t.Errorf("ClassFare(%d, %t) = %d; want %d", tt.age, tt.isStudent, got, tt.want)
		}
	}
}
