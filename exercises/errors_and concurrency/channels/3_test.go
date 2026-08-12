package main

import "testing"

func TestEvenSum(t *testing.T) {
	if got := EvenSum(10); got != 30 {
		t.Errorf("EvenSum = %d; want 30", got)
	}
}
