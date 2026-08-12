package main

import (
	"reflect"
	"testing"
)

func TestLoopCleanup(t *testing.T) {
	got := LoopCleanup(3)
	want := []int{2, 1, 0}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("LoopCleanup(3) = %v; want %v", got, want)
	}
}
