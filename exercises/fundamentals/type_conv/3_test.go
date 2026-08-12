package main

import "testing"

func TestToPixels(t *testing.T) {
	if got := ToPixels(1.5, 300); got != 450 {
		t.Errorf("ToPixels(1.5, 300) = %d; want 450", got)
	}
	if got := ToPixels(0.5, 72); got != 36 {
		t.Errorf("ToPixels(0.5, 72) = %d; want 36", got)
	}
}
