package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestDivide(t *testing.T) {
	tests := []struct {
		name        string
		a, b        float64
		wantResult  float64
		wantErrMsg  string
		expectError bool
	}{
		{"normal division", 10.0, 2.0, 5.0, "", false},
		{"division by zero", 5.0, 0.0, 0.0, "cannot divide by zero", true},
		{"negative result", -6.0, 3.0, -2.0, "", false},
		{"zero numerator", 0.0, 5.0, 0.0, "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, gotErr := Divide(tt.a, tt.b)
			if tt.expectError {
				if gotErr == nil {
					t.Errorf("Divide(%f, %f) expected error, got nil", tt.a, tt.b)
				} else if gotErr.Error() != tt.wantErrMsg {
					t.Errorf("Divide(%f, %f) expected error message %q, got %q", tt.a, tt.b, tt.wantErrMsg, gotErr.Error())
				}
			} else {
				if gotErr != nil {
					t.Errorf("Divide(%f, %f) expected no error, got: %v", tt.a, tt.b, gotErr)
				}
				if gotResult != tt.wantResult {
					t.Errorf("Divide(%f, %f) expected %f, got %f", tt.a, tt.b, tt.wantResult, gotResult)
				}
			}
		})
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
		"Result: 5.00",
		"Error: cannot divide by zero",
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
