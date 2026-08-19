package main

import (
	"os"
	"regexp"
	"testing"
)

func TestBlueValue(t *testing.T) {
	if got := BlueValue(); got != 3 {
		t.Errorf("BlueValue() = %d; want 3", got)
	}
	data, err := os.ReadFile("2.go")
	if err != nil {
		t.Fatalf("read 2.go: %v", err)
	}
	if !regexp.MustCompile(`(?s)const\s*\([^)]*iota[^)]*\)`).Match(data) {
		t.Errorf("const block must use iota")
	}
	for _, name := range []string{"Red", "Green", "Blue", "Alpha"} {
		if !regexp.MustCompile(`(?s)const\s*\([^)]*\b` + name + `\b[^)]*\)`).Match(data) {
			t.Errorf("const block must define %s", name)
		}
	}
}