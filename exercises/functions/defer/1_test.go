package main

import (
	"reflect"
	"testing"
)

func TestProcessLog(t *testing.T) {
	got := ProcessLog()
	want := []string{"start", "middle", "end"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ProcessLog = %v; want %v", got, want)
	}
}
