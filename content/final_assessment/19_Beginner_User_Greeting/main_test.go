package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func runWithInput(t *testing.T, input string) string {
	t.Helper()
	oldStdin := os.Stdin
	oldStdout := os.Stdout

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("stdin pipe: %v", err)
	}
	if _, err := w.WriteString(input); err != nil {
		t.Fatalf("write stdin: %v", err)
	}
	if err := w.Close(); err != nil {
		t.Fatalf("close stdin: %v", err)
	}
	os.Stdin = r

	outR, outW, err := os.Pipe()
	if err != nil {
		t.Fatalf("stdout pipe: %v", err)
	}
	os.Stdout = outW

	main()

	outW.Close()
	os.Stdin = oldStdin
	os.Stdout = oldStdout
	r.Close()

	var buf bytes.Buffer
	io.Copy(&buf, outR)
	outR.Close()
	return strings.TrimSpace(buf.String())
}

func TestUserGreeting(t *testing.T) {
	if got := runWithInput(t, "Gopher\n"); got != "Hello, Gopher!" {
		t.Errorf("output = %q; want %q", got, "Hello, Gopher!")
	}
}

func TestUserGreetingAda(t *testing.T) {
	if got := runWithInput(t, "Ada\n"); got != "Hello, Ada!" {
		t.Errorf("output = %q; want %q", got, "Hello, Ada!")
	}
}