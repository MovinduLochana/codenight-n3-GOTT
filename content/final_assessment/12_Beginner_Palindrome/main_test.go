package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"Racecar", true},
		{"golang", false},
		{"Aba", true},
		{"hello", false},
		{"a", true},
		{"", true},
	}

	for _, tt := range tests {
		got := IsPalindrome(tt.input)
		if got != tt.want {
			t.Errorf("IsPalindrome(%q) = %t; want %t", tt.input, got, tt.want)
		}
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
	lines := strings.Split(output, "\n")
	for i := range lines {
		lines[i] = strings.TrimSpace(lines[i])
	}

	expected := []string{
		"Racecar: true",
		"golang: false",
	}

	if len(lines) != len(expected) {
		t.Errorf("Expected %d lines of output, got %d. Output:\n%s", len(expected), len(lines), output)
		return
	}

	for i, line := range lines {
		if line != expected[i] {
			t.Errorf("At line %d: expected %q, got %q", i+1, expected[i], line)
		}
	}
}
