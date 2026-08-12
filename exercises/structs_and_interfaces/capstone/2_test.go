package main

import "testing"

func TestTotalPrice(t *testing.T) {
	items := []Pricer{Product{Name: "Coffee", BasePrice: 4.50}, Product{Name: "Tea", BasePrice: 2.50}}
	if got := TotalPrice(items); got != 7.0 {
		t.Errorf("TotalPrice = %f; want 7.0", got)
	}
}
