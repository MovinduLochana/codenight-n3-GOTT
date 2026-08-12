package main

import (
	"reflect"
	"testing"
)

func TestInsertAt(t *testing.T) {
	got := InsertAt([]int{10, 20, 30}, 1, 99)
	want := []int{10, 99, 20, 30}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("InsertAt = %v; want %v", got, want)
	}
}
