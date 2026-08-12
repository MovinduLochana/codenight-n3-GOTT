package main

import "testing"

func TestIsEven(t *testing.T) {
	if !IsEven(4) {
		t.Errorf("IsEven(4) = false; want true")
	}
	if IsEven(7) {
		t.Errorf("IsEven(7) = true; want false")
	}
}
