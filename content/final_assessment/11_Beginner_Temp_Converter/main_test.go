package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

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

	// Normalize degree symbols just in case
	output = strings.ReplaceAll(output, "°", "")
	expected := "98.6F is 37.0C"
	if strings.ReplaceAll(output, "°", "") != expected {
		t.Errorf("Expected output %q (ignoring degree symbol formatting), got %q", expected, output)
	}
}
