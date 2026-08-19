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

func TestInsertAt(t *testing.T) {
	tests := []struct {
		name string
		s    []int
		i, v int
		want []int
	}{
		{"middle", []int{10, 20, 30}, 1, 99, []int{10, 99, 20, 30}},
		{"first", []int{1, 2}, 0, 0, []int{0, 1, 2}},
		{"last", []int{1, 2}, 2, 5, []int{1, 2, 5}},
		{"empty", []int{}, 0, 7, []int{7}},
		{"single", []int{7}, 0, 3, []int{3, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertAt(tt.s, tt.i, tt.v); !sameInts(got, tt.want) {
				t.Errorf("InsertAt(%v, %d, %d) = %v; want %v", tt.s, tt.i, tt.v, got, tt.want)
			}
		})
	}

	orig := []int{1, 2, 3}
	got := InsertAt(orig, 1, 99)
	got[0] = 88
	if orig[0] == 88 {
		t.Errorf("InsertAt must return a new slice, not share backing array with the input")
	}
}