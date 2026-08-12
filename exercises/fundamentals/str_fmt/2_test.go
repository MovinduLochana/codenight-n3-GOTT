package main

import (
	"strings"
	"testing"
)

func TestReport(t *testing.T) {
	got := Report("Alice", 3, 150.0)
	if !strings.Contains(got, "Alice") || !strings.Contains(got, "3") {
		t.Errorf("Report invalid: %q", got)
	}
}
