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

func TestRemoveAt(t *testing.T) {
	tests := []struct {
		name string
		s    []int
		i    int
		want []int
	}{
		{"middle", []int{10, 20, 30}, 1, []int{10, 30}},
		{"first", []int{5, 6, 7, 8}, 0, []int{6, 7, 8}},
		{"last", []int{1, 2, 3}, 2, []int{1, 2}},
		{"single", []int{9}, 0, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveAt(tt.s, tt.i); !sameInts(got, tt.want) {
				t.Errorf("RemoveAt(%v, %d) = %v; want %v", tt.s, tt.i, got, tt.want)
			}
		})
	}

	orig := []int{1, 2, 3}
	got := RemoveAt(orig, 1)
	got[0] = 99
	if orig[0] == 99 || orig[2] == 99 {
		t.Errorf("RemoveAt must return a new slice, not mutate or share with the input")
	}
}