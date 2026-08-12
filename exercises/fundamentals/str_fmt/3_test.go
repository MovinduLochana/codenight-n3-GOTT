package main

import "testing"

func TestRow(t *testing.T) {
	got := Row("Coffee", 4.50)
	if len(got) == 0 {
		t.Errorf("Row returned empty string")
	}
}
