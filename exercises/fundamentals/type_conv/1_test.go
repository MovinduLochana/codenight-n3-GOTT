package main

import "testing"

func TestAverage(t *testing.T) {
	got := Average(3, 4)
	if got != 3.5 {
		t.Errorf("Average(3, 4) = %f; want 3.5", got)
	}
}
