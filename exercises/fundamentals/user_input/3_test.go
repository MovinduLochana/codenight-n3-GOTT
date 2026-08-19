package main

import (
	"os"
	"testing"
)

func withInput(t *testing.T, input string, fn func()) {
	t.Helper()
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("pipe: %v", err)
	}
	if _, err := w.WriteString(input); err != nil {
		t.Fatalf("write stdin: %v", err)
	}
	if err := w.Close(); err != nil {
		t.Fatalf("close pipe: %v", err)
	}
	os.Stdin = r
	defer r.Close()

	fn()
}

func TestAverageUntilZero(t *testing.T) {
	withInput(t, "10 20 30 0\n", func() {
		if got := AverageUntilZero(); got != 20 {
			t.Errorf("AverageUntilZero() = %v; want 20", got)
		}
	})
	withInput(t, "5 0\n", func() {
		if got := AverageUntilZero(); got != 5 {
			t.Errorf("AverageUntilZero() = %v; want 5", got)
		}
	})
}