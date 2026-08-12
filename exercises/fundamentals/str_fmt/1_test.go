package main

import "testing"

func TestReceipt(t *testing.T) {
	expected := "Item: Coffee, Qty: 2, Price: $4.50"
	if got := Receipt("Coffee", 2, 4.50); got != expected {
		t.Errorf("Receipt = %q; want %q", got, expected)
	}
}
