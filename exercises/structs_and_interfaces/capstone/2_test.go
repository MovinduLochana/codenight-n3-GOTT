package main

import (
	"math"
	"testing"
)

func TestTotalPrice(t *testing.T) {
	p := Product{Name: "Coffee", BasePrice: 4.50}
	if got := p.Price(); got != 4.50 {
		t.Errorf("Product.Price() = %v; want 4.50", got)
	}

	items := []Pricer{
		Product{Name: "Coffee", BasePrice: 4.50},
		Product{Name: "Cake", BasePrice: 6.00},
	}
	if got := TotalPrice(items); math.Abs(got-10.50) > 1e-9 {
		t.Errorf("TotalPrice(items) = %v; want 10.50", got)
	}

	if got := TotalPrice([]Pricer{}); got != 0 {
		t.Errorf("TotalPrice(empty) = %v; want 0", got)
	}
}