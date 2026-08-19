package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestRow(t *testing.T) {
	tests := []struct {
		item  string
		price float64
	}{
		{"Coffee", 450.5},
		{"Short", 25.0},
		{"Espresso", 9.99},
	}
	for _, tt := range tests {
		got := Row(tt.item, tt.price)
		if !strings.HasPrefix(got, tt.item) {
			t.Errorf("Row(%q, %v) = %q; want it to start with the item name", tt.item, tt.price, got)
		}
		if !strings.HasSuffix(got, fmt.Sprintf("%.2f", tt.price)) {
			t.Errorf("Row(%q, %v) = %q; want it to end with the 2-decimal price", tt.item, tt.price, got)
		}
		if len(got) <= len(tt.item)+len(fmt.Sprintf("%.2f", tt.price)) {
			t.Errorf("Row(%q, %v) = %q; expected padding between name and price", tt.item, tt.price, got)
		}
	}
}