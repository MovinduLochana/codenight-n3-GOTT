package main

import (
	"reflect"
	"testing"
)

func TestRemoveAt(t *testing.T) {
	got := RemoveAt([]int{10, 20, 30}, 1)
	want := []int{10, 30}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("RemoveAt = %v; want %v", got, want)
	}
}
