package main

import "testing"

func sameInts(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestCheckout(t *testing.T) {
	result, err := Checkout([]string{"Coffee", "Tea"}, []int{2, 3})
	if err != nil {
		t.Errorf("Checkout(valid) error = %v; want nil", err)
	}
	if want := []int{4, 6}; !sameInts(result, want) {
		t.Errorf("Checkout(valid) = %v; want %v", result, want)
	}

	result, err = Checkout([]string{"Coffee"}, []int{0})
	if err == nil {
		t.Errorf("Checkout(invalid) returned nil error; want non-nil")
	}
	if result != nil {
		t.Errorf("Checkout(invalid) result = %v; want nil", result)
	}

	result, err = Checkout([]string{"Tea"}, []int{1})
	if err != nil {
		t.Errorf("Checkout(single) error = %v; want nil", err)
	}
	if want := []int{2}; !sameInts(result, want) {
		t.Errorf("Checkout(single) = %v; want %v", result, want)
	}

	result, err = Checkout([]string{}, []int{})
	if err != nil {
		t.Errorf("Checkout(empty) error = %v; want nil", err)
	}
	if want := []int{}; !sameInts(result, want) {
		t.Errorf("Checkout(empty) = %v; want %v", result, want)
	}
}