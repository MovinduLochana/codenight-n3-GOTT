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

func TestGreet(t *testing.T) {
	withInput(t, "Gopher\n", func() {
		if got := Greet(); got != "Hello, Gopher!" {
			t.Errorf("Greet() = %q; want %q", got, "Hello, Gopher!")
		}
	})
	withInput(t, "Ada\n", func() {
		if got := Greet(); got != "Hello, Ada!" {
			t.Errorf("Greet() = %q; want %q", got, "Hello, Ada!")
		}
	})
}