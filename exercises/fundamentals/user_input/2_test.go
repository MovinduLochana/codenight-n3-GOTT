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

func TestSum(t *testing.T) {
	withInput(t, "3 4\n", func() {
		if got := Sum(); got != 7 {
			t.Errorf("Sum() = %d; want 7", got)
		}
	})
	withInput(t, "10 20\n", func() {
		if got := Sum(); got != 30 {
			t.Errorf("Sum() = %d; want 30", got)
		}
	})
}