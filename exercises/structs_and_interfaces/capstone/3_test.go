package main

import "testing"

func TestDealPrice(t *testing.T) {
	items := []Pricer{
		Product{Name: "Coffee", BasePrice: 10.0},
		Deal{Product: Product{Name: "Tea", BasePrice: 10.0}, Discount: 0.2},
	}
	if got := TotalPrice(items); got != 18.0 {
		t.Errorf("TotalPrice = %f; want 18.0", got)
	}
}
