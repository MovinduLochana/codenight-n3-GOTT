package main

import (
	"reflect"
	"testing"
)

func TestCut(t *testing.T) {
	got := Cut([]int{10, 20, 30, 40, 50}, 1, 3)
	want := []int{20, 30}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Cut = %v; want %v", got, want)
	}
}
