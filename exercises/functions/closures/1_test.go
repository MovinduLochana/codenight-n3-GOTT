package main

import "testing"

func TestMakeCounter(t *testing.T) {
	c := MakeCounter()
	if c() != 1 || c() != 2 || c() != 3 {
		t.Errorf("MakeCounter increment failed")
	}
}
