package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestCountLines(t *testing.T) {
	// Create a temp file for testing
	tmpFile, err := os.CreateTemp("", "test_count_lines_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	content := "Line 1\nLine 2\nLine 3\nLine 4\n"
	if _, err := tmpFile.WriteString(content); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tmpFile.Close()

	count, err := CountLines(tmpFile.Name())
	if err != nil {
		t.Fatalf("CountLines failed: %v", err)
	}
	if count != 4 {
		t.Errorf("Expected 4 lines, got %d", count)
	}

	// Test non-existent file error
	_, err = CountLines("does_not_exist_xyz.txt")
	if err == nil {
		t.Error("Expected error when counting lines of non-existent file, got nil")
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

	// Since main uses "sample.txt", and sample.txt has exactly 3 lines:
	expected := "Line count: 3"
	if output != expected {
		t.Errorf("Expected output %q, got %q. Ensure 'sample.txt' exists and contains 3 lines.", expected, output)
	}
}
