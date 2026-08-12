package main

import "testing"

func TestIsLeapYear(t *testing.T) {
	if !IsLeapYear(2024) {
		t.Errorf("IsLeapYear(2024) = false; want true")
	}
	if IsLeapYear(1900) {
		t.Errorf("IsLeapYear(1900) = true; want false")
	}
	if !IsLeapYear(2000) {
		t.Errorf("IsLeapYear(2000) = false; want true")
	}
}
