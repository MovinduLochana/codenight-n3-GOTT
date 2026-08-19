package main

import "testing"

func TestReceipt(t *testing.T) {
	tests := []struct {
		item  string
		qty   int
		price float64
		want  string
	}{
		{"Coffee", 2, 4.50, "Item: Coffee, Qty: 2, Price: $4.50"},
		{"Tea", 1, 3.00, "Item: Tea, Qty: 1, Price: $3.00"},
		{"Cake", 3, 10.0, "Item: Cake, Qty: 3, Price: $10.00"},
	}
	for _, tt := range tests {
		if got := Receipt(tt.item, tt.qty, tt.price); got != tt.want {
			t.Errorf("Receipt(%q, %d, %v) = %q; want %q", tt.item, tt.qty, tt.price, got, tt.want)
		}
	}
}