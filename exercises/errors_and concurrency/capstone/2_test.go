package main

import "testing"

func TestCheckoutTotals(t *testing.T) {
	if got := CheckoutTotals(); got != 60 {
		t.Errorf("CheckoutTotals = %d; want 60", got)
	}
}
