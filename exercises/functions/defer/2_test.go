package main

import (
	"reflect"
	"testing"
)

func TestStackLog(t *testing.T) {
	want := []string{"start", "A", "B"}
	if got := StackLog(); !reflect.DeepEqual(got, want) {
		t.Errorf("StackLog() = %v; want %v", got, want)
	}
}