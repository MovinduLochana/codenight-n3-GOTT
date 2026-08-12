package main

import "testing"

func TestDivide(t *testing.T) {
	res, err := Divide(10, 2)
	if err != nil || res != 5 {
		t.Errorf("Divide(10, 2) = %d, %v; want 5, nil", res, err)
	}
	_, err = Divide(5, 0)
	if err == nil {
		t.Errorf("Divide(5, 0) expected error")
	}
}
