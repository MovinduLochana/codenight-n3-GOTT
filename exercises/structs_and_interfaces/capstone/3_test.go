package main

import (
	"math"
	"testing"
)

func TestDealPrice(t *testing.T) {
	d := Deal{Product: Product{Name: "Cake", BasePrice: 10.00}, Discount: 0.20}
	if got := d.Price(); math.Abs(got-8.00) > 1e-9 {
		t.Errorf("Deal.Price() = %v; want 8.00 (10.00 * (1 - 0.20))", got)
	}

	items := []Pricer{
		Product{Name: "Coffee", BasePrice: 4.00},
		Deal{Product: Product{Name: "Cake", BasePrice: 10.00}, Discount: 0.20},
	}
	if got := TotalPrice(items); math.Abs(got-12.00) > 1e-9 {
		t.Errorf("TotalPrice(items) = %v; want 12.00", got)
	}

	fullPrice := Deal{Product: Product{Name: "X", BasePrice: 5.00}, Discount: 0}
	if got := fullPrice.Price(); got != 5.00 {
		t.Errorf("Deal with 0 discount = %v; want 5.00", got)
	}
}