package main

import (
	"reflect"
	"testing"
)

func TestSquares(t *testing.T) {
	got := Squares([]int{4, 1, 3})
	want := []int{1, 9, 16}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Squares = %v; want %v", got, want)
	}
}
