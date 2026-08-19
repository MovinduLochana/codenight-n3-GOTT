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

func TestProducer(t *testing.T) {
	tests := []struct {
		n    int
		want []int
	}{
		{3, []int{1, 2, 3}},
		{1, []int{1}},
		{0, []int{}},
		{5, []int{1, 2, 3, 4, 5}},
	}
	for _, tt := range tests {
		ch := make(chan int, tt.n)
		Producer(ch, tt.n)
		var got []int
		for v := range ch {
			got = append(got, v)
		}
		if !sameInts(got, tt.want) {
			t.Errorf("Producer(ch, %d) sent %v; want %v", tt.n, got, tt.want)
		}
	}
}