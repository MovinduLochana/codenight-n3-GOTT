package main

import "testing"

func TestMoney(t *testing.T) {
	if got := Money(12.5).Format(); got != 12.5 {
		t.Errorf("Money(12.5).Format() = %v; want 12.5", got)
	}
	if got := Money(0).Format(); got != 0 {
		t.Errorf("Money(0).Format() = %v; want 0", got)
	}

	if got := AddCurrency(Money(1.25), Money(2.50)); got != 3.75 {
		t.Errorf("AddCurrency(1.25, 2.50) = %v; want 3.75", got)
	}
	if got := AddCurrency(Money(0), Money(0)); got != 0 {
		t.Errorf("AddCurrency(0, 0) = %v; want 0", got)
	}
}