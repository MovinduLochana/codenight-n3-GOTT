package main

import (
	"reflect"
	"testing"
)

func TestSquaresPipeline(t *testing.T) {
	got := SquaresPipeline(4)
	want := []int{1, 4, 9, 16}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("SquaresPipeline = %v; want %v", got, want)
	}
}
