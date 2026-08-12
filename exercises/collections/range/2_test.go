package main

import "testing"

func TestContains(t *testing.T) {
	if !Contains([]int{1, 2, 3}, 2) {
		t.Errorf("Contains(2) = false; want true")
	}
	if Contains([]int{1, 2, 3}, 9) {
		t.Errorf("Contains(9) = true; want false")
	}
}
