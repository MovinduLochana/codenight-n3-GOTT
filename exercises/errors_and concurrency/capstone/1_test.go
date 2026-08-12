package main

import "testing"

func TestValidate(t *testing.T) {
	if err := Validate("Coffee", 2); err != nil {
		t.Errorf("Validate valid item returned error: %v", err)
	}
	if err := Validate("Tea", 0); err == nil {
		t.Errorf("Validate invalid item expected error")
	}
}
