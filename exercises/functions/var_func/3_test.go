package main

import "testing"

func TestMean(t *testing.T) {
	if got := Mean(90, 86, 88); got != 88.0 {
		t.Errorf("Mean = %f; want 88.0", got)
	}
}
