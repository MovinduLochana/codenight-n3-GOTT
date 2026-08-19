package main

import "testing"

func TestClassFare(t *testing.T) {
	tests := []struct {
		age       int
		isStudent bool
		want      int
	}{
		{8, false, 25},   // under 12
		{70, false, 25},  // 65+
		{20, true, 30},   // student
		{20, false, 50},  // everyone else
		{10, true, 25},   // under 12 wins over student
		{11, true, 25},   // under 12
		{12, true, 30},   // exactly 12 is a student
		{65, true, 25},   // 65+ wins over student
		{66, true, 25},   // 65+
		{12, false, 50},  // 12 not a student
		{64, true, 30},   // student under 65
	}
	for _, tt := range tests {
		if got := ClassFare(tt.age, tt.isStudent); got != tt.want {
			t.Errorf("ClassFare(%d, %t) = %d; want %d", tt.age, tt.isStudent, got, tt.want)
		}
	}
}