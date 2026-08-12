package main

import "testing"

func TestBlueValue(t *testing.T) {
	if got := BlueValue(); got != 3 {
		t.Errorf("BlueValue() = %d; want 3", got)
	}
}
