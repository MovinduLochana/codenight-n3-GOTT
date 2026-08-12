package main

import "testing"

func TestDivMod(t *testing.T) {
	q, r, err := DivMod(10, 3)
	if err != nil || q != 3 || r != 1 {
		t.Errorf("DivMod(10, 3) = %d, %d, %v; want 3, 1, nil", q, r, err)
	}
	_, _, err = DivMod(5, 0)
	if err == nil {
		t.Errorf("DivMod(5, 0) expected error")
	}
}
