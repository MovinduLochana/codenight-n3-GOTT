package main

import "testing"

func TestAllPermissions(t *testing.T) {
	if got := AllPermissions(); got != 7 {
		t.Errorf("AllPermissions() = %d; want 7", got)
	}
}
