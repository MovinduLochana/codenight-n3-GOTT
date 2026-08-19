package main

import "testing"

func TestLabel(t *testing.T) {
	tests := []struct {
		p    Product
		want string
	}{
		{Product{Name: "Coffee", Price: 4.50}, "Coffee: $4.50"},
		{Product{Name: "Tea", Price: 3.00}, "Tea: $3.00"},
		{Product{Name: "Free", Price: 0}, "Free: $0.00"},
		{Product{Name: "Big", Price: 12.5}, "Big: $12.50"},
	}
	for _, tt := range tests {
		if got := tt.p.Label(); got != tt.want {
			t.Errorf("Product{Name:%q, Price:%v}.Label() = %q; want %q", tt.p.Name, tt.p.Price, got, tt.want)
		}
	}
}