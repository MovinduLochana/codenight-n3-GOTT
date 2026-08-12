package main

import "testing"

func TestLabel(t *testing.T) {
	p := Product{Name: "Coffee", Price: 4.50}
	if got := p.Label(); got != "Coffee: $4.50" {
		t.Errorf("Label = %q; want Coffee: $4.50", got)
	}
}
