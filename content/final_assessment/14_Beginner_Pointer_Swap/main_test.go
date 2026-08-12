package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestSwap(t *testing.T) {
	a, b := 100, 200
	Swap(&a, &b)
	if a != 200 || b != 100 {
		t.Errorf("Expected swap(100, 200) to result in (200, 100), got (%d, %d)", a, b)
	}
}

func TestMainOutput(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := strings.TrimSpace(buf.String())

	expected := "x: 20, y: 10"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
