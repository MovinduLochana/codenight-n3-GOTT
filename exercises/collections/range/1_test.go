package main

import "testing"

func TestSum(t *testing.T) {
	if got := Sum([]int{1, 2, 3, 4}); got != 10 {
		t.Errorf("Sum = %d; want 10", got)
	}
	if got := Sum([]int{}); got != 0 {
		t.Errorf("Sum = %d; want 0", got)
	}
}
