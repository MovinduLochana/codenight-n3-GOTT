package main

import (
	"os"
	"regexp"
	"testing"
)

func TestPriorityValue(t *testing.T) {
	if got := PriorityValue(); got != 1 {
		t.Errorf("PriorityValue() = %d; want 1", got)
	}
	data, err := os.ReadFile("1.go")
	if err != nil {
		t.Fatalf("read 1.go: %v", err)
	}
	if !regexp.MustCompile(`(?s)const\s*\([^)]*iota[^)]*\)`).Match(data) {
		t.Errorf("const block must use iota")
	}
	for _, name := range []string{"Low", "Medium", "High"} {
		if !regexp.MustCompile(`(?s)const\s*\([^)]*\b` + name + `\b[^)]*\)`).Match(data) {
			t.Errorf("const block must define %s", name)
		}
	}
}