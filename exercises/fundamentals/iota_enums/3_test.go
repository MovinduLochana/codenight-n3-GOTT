package main

import (
	"os"
	"regexp"
	"testing"
)

func TestAllPermissions(t *testing.T) {
	if got := AllPermissions(); got != 7 {
		t.Errorf("AllPermissions() = %d; want 7", got)
	}
	data, err := os.ReadFile("3.go")
	if err != nil {
		t.Fatalf("read 3.go: %v", err)
	}
	if !regexp.MustCompile(`(?s)const\s*\([^)]*1\s*<<\s*iota[^)]*\)`).Match(data) {
		t.Errorf("const block must use `1 << iota`")
	}
	for _, name := range []string{"Read", "Write", "Execute"} {
		if !regexp.MustCompile(`(?s)const\s*\([^)]*\b` + name + `\b[^)]*\)`).Match(data) {
			t.Errorf("const block must define %s", name)
		}
	}
}