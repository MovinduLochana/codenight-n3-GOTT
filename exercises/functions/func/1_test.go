package main

import "testing"

func TestDivide(t *testing.T) {
	q, err := Divide(10, 2)
	if err != nil {
		t.Errorf("Divide(10, 2) returned error %v; want nil", err)
	}
	if q != 5 {
		t.Errorf("Divide(10, 2) = %d; want 5", q)
	}

	if _, err := Divide(5, 0); err == nil {
		t.Errorf("Divide(5, 0) returned nil error; want a non-nil error")
	}
}