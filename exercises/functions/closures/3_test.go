package main

import "testing"

func TestMakeBank(t *testing.T) {
	bank := MakeBank(100)
	steps := []struct {
		amount float64
		want   float64
	}{
		{50, 150},
		{-20, 150},
		{0, 150},
		{25, 175},
		{-100, 175},
	}
	for _, s := range steps {
		if got := bank(s.amount); got != s.want {
			t.Errorf("bank(%v) = %v; want %v", s.amount, got, s.want)
		}
	}
}