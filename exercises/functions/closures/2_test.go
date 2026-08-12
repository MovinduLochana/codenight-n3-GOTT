package main

import "testing"

func TestMakeAdder(t *testing.T) {
	add5 := MakeAdder(5)
	if got := add5(10); got != 15 {
		t.Errorf("MakeAdder(5)(10) = %d; want 15", got)
	}
}
