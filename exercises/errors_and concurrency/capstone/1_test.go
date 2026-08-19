package main

import "testing"

func TestValidate(t *testing.T) {
	if err := Validate("Coffee", 2); err != nil {
		t.Errorf("Validate(\"Coffee\", 2) = %v; want nil", err)
	}
	if err := Validate("Coffee", 1); err != nil {
		t.Errorf("Validate(\"Coffee\", 1) = %v; want nil", err)
	}
	if err := Validate("Tea", 0); err == nil {
		t.Errorf("Validate(\"Tea\", 0) returned nil; want error")
	}
	if err := Validate("Tea", -3); err == nil {
		t.Errorf("Validate(\"Tea\", -3) returned nil; want error")
	}
}