package main

import "testing"

func TestMoney(t *testing.T) {
	got := AddCurrency(Money(10.5), Money(4.5))
	if got != Money(15.0) {
		t.Errorf("AddCurrency = %v; want 15.0", got)
	}
}
