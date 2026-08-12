package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestStack(t *testing.T) {
	var s Stack

	// Test Pop on empty stack
	val, ok := s.Pop()
	if val != "" || ok {
		t.Errorf("Expected empty stack Pop to return (\"\", false), got (%q, %t)", val, ok)
	}

	// Test Push
	s.Push("A")
	if len(s.elements) != 1 || s.elements[0] != "A" {
		t.Errorf("Expected stack elements to be [A] after Push(\"A\"), got %v", s.elements)
	}

	s.Push("B")
	s.Push("C")

	// Test Pop (LIFO order)
	val, ok = s.Pop()
	if val != "C" || !ok {
		t.Errorf("Expected first Pop to return (\"C\", true), got (%q, %t)", val, ok)
	}

	val, ok = s.Pop()
	if val != "B" || !ok {
		t.Errorf("Expected second Pop to return (\"B\", true), got (%q, %t)", val, ok)
	}

	val, ok = s.Pop()
	if val != "A" || !ok {
		t.Errorf("Expected third Pop to return (\"A\", true), got (%q, %t)", val, ok)
	}

	// Stack should be empty again
	val, ok = s.Pop()
	if val != "" || ok {
		t.Errorf("Expected stack to be empty after popping all elements, got (%q, %t)", val, ok)
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
		"Beginning transaction...",
		"Transaction steps registered.",
		"Rolling back or auditing transaction steps:",
		"- COMMIT",
		"- WRITE_TEMP",
		"- START",
	}

	if len(lines) != len(expected) {
		t.Errorf("Expected %d lines of output, got %d. Output was:\n%s", len(expected), len(lines), output)
		return
	}

	for i, line := range lines {
		if line != expected[i] {
			t.Errorf("At line %d: expected %q, got %q", i+1, expected[i], line)
		}
	}
}
