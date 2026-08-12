package main

import (
	"reflect"
	"testing"
)

func TestStackLog(t *testing.T) {
	got := StackLog()
	want := []string{"start", "A", "B"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("StackLog = %v; want %v", got, want)
	}
}
