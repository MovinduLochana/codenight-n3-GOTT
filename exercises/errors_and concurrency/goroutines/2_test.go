package main

import "testing"

func TestConcurrentSum(t *testing.T) {
	if got := ConcurrentSum([]int{1, 2, 3, 4}); got != 10 {
		t.Errorf("ConcurrentSum = %d; want 10", got)
	}
}
