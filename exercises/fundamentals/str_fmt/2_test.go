package main

import (
	"strconv"
	"strings"
	"testing"
)

func TestReport(t *testing.T) {
	got := Report("Nimal", 12, 3450.5)
	checks := []struct {
		name  string
		found bool
	}{
		{"name", strings.Contains(got, "Nimal")},
		{"orders", strings.Contains(got, strconv.Itoa(12))},
		{"revenue", strings.Contains(got, "3450.50")},
		{"newline", strings.Contains(got, "\n")},
	}
	for _, c := range checks {
		if !c.found {
			t.Errorf("Report output missing %s: %q", c.name, got)
		}
	}
}