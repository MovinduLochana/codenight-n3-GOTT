package main

import "testing"

func TestPriorityValue(t *testing.T) {
	if got := PriorityValue(); got != 1 {
		t.Errorf("PriorityValue() = %d; want 1", got)
	}
}
