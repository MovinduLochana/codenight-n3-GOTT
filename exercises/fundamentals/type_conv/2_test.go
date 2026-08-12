package main

import "testing"

func TestCelsius(t *testing.T) {
	if got := Celsius(68); got != 20 {
		t.Errorf("Celsius(68) = %f; want 20", got)
	}
	if got := Celsius(32); got != 0 {
		t.Errorf("Celsius(32) = %f; want 0", got)
	}
}
