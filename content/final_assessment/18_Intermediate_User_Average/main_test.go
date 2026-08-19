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

func TestUserAverage(t *testing.T) {
	if got := runWithInput(t, "10 20 30 0\n"); got != "Average: 20" {
		t.Errorf("output = %q; want %q", got, "Average: 20")
	}
}

func TestUserAverageSingle(t *testing.T) {
	if got := runWithInput(t, "7 0\n"); got != "Average: 7" {
		t.Errorf("output = %q; want %q", got, "Average: 7")
	}
}

func TestUserAverageImmediateZero(t *testing.T) {
	if got := runWithInput(t, "0\n"); got != "Average: 0" {
		t.Errorf("output = %q; want %q", got, "Average: 0")
	}
}