package main

import "testing"

func TestMinMax(t *testing.T) {
	min, max := MinMax([]int{3, 1, 4, 1, 5})
	if min != 1 || max != 5 {
		t.Errorf("MinMax = %d, %d; want 1, 5", min, max)
	}
}
