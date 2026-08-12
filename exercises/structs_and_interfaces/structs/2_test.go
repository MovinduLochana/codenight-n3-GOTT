package main

import "testing"

func TestVolume(t *testing.T) {
	b := Box{Length: 2, Width: 3, Height: 4}
	if got := Volume(b); got != 24.0 {
		t.Errorf("Volume = %f; want 24.0", got)
	}
}
