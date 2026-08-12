package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestMainOutput(t *testing.T) {
	// Keep backup of original stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Run main
	main()

	// Close write and restore stdout
	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := strings.TrimSpace(buf.String())

	expected := "Average Score: 78.33"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
