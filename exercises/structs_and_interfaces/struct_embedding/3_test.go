package main

import "testing"

func TestTotalAreaEmbedding(t *testing.T) {
	shapes := []Shape{Square{Side: 3}}
	if got := TotalArea(shapes); got != 9.0 {
		t.Errorf("TotalArea = %f; want 9.0", got)
	}
}
