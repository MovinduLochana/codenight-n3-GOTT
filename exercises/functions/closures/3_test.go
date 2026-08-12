package main

import "testing"

func TestMakeBank(t *testing.T) {
	bank := MakeBank(100)
	if got := bank(50); got != 150 {
		t.Errorf("bank(50) = %f; want 150", got)
	}
	if got := bank(-20); got != 150 {
		t.Errorf("bank(-20) = %f; want 150", got)
	}
	if got := bank(25); got != 175 {
		t.Errorf("bank(25) = %f; want 175", got)
	}
}
