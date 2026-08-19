package main

import (
	"reflect"
	"testing"
)

func TestProcessLog(t *testing.T) {
	want := []string{"start", "middle", "end"}
	if got := ProcessLog(); !reflect.DeepEqual(got, want) {
		t.Errorf("ProcessLog() = %v; want %v", got, want)
	}
}