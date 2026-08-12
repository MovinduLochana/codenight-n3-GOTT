package main

import "testing"

func TestIsBetween(t *testing.T) {
	if !IsBetween(5, 1, 10) {
		t.Errorf("IsBetween(5, 1, 10) = false; want true")
	}
	if IsBetween(0, 1, 10) {
		t.Errorf("IsBetween(0, 1, 10) = true; want false")
	}
}
