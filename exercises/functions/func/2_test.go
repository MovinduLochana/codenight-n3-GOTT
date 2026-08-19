package main

import "testing"

func TestDivMod(t *testing.T) {
	q, r, err := DivMod(10, 3)
	if err != nil {
		t.Errorf("DivMod(10, 3) returned error %v; want nil", err)
	}
	if q != 3 || r != 1 {
		t.Errorf("DivMod(10, 3) = (%d, %d); want (3, 1)", q, r)
	}

	q, r, err = DivMod(10, 0)
	if err == nil {
		t.Errorf("DivMod(10, 0) returned nil error; want a non-nil error")
	}
	if q != 0 || r != 0 {
		t.Errorf("DivMod(10, 0) = (%d, %d); want (0, 0)", q, r)
	}
}