package main

import (
	"reflect"
	"testing"
)

func TestIndexAll(t *testing.T) {
	got := IndexAll([]int{1, 2, 1, 4, 1}, 1)
	want := []int{0, 2, 4}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("IndexAll = %v; want %v", got, want)
	}
}
