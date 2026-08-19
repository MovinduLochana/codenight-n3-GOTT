package main

import "testing"

func sameInts(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestTail(t *testing.T) {
	tests := []struct {
		name string
		s    []int
		n    int
		want []int
	}{
		{"basic", []int{1, 2, 3, 4, 5}, 2, []int{4, 5}},
		{"n larger than len", []int{1, 2}, 5, []int{1, 2}},
		{"zero", []int{1, 2, 3}, 0, []int{}},
		{"exact", []int{7}, 1, []int{7}},
		{"empty", []int{}, 3, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Tail(tt.s, tt.n); !sameInts(got, tt.want) {
				t.Errorf("Tail(%v, %d) = %v; want %v", tt.s, tt.n, got, tt.want)
			}
		})
	}

	orig := []int{1, 2, 3, 4, 5}
	got := Tail(orig, 2)
	got[0] = 99
	if orig[3] == 99 {
		t.Errorf("Tail must return an independent copy, not share backing array")
	}
}