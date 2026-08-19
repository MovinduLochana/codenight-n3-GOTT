package main

import "testing"

func TestMakeCounter(t *testing.T) {
	counter := MakeCounter()
	for i, want := range []int{1, 2, 3, 4, 5} {
		if got := counter(); got != want {
			t.Errorf("counter() call %d = %d; want %d", i+1, got, want)
		}
	}

	other := MakeCounter()
	if got := other(); got != 1 {
		t.Errorf("independent counter first call = %d; want 1", got)
	}
	if got := counter(); got != 6 {
		t.Errorf("original counter after other created = %d; want 6 (state must not be shared)", got)
	}
}