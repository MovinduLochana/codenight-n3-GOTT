package main

import "testing"

func TestAddItem(t *testing.T) {
	c := Cart{}
	p := Product{Name: "Coffee", Price: 4.50}
	c2 := AddItem(c, p)
	if len(c2.Items) != 1 || c2.Items[0].Name != "Coffee" {
		t.Errorf("AddItem failed")
	}
}
