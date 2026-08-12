package main

import (
	"bytes"
	"errors"
	"io"
	"os"
	"strings"
	"testing"
)

func TestCustomError(t *testing.T) {
	// Test normal case (no error)
	err := VerifyAge(25)
	if err != nil {
		t.Errorf("VerifyAge(25) returned unexpected error: %v", err)
	}

	// Test invalid negative age
	err = VerifyAge(-1)
	if err == nil {
		t.Fatal("VerifyAge(-1) expected error, got nil")
	}

	// Verify error implements ErrInvalidAge and has the right value
	var invalidAgeErr *ErrInvalidAge
	if !errors.As(err, &invalidAgeErr) {
		t.Errorf("Expected error to be of type *ErrInvalidAge, got %T", err)
	} else if invalidAgeErr.Age != -1 {
		t.Errorf("Expected ErrInvalidAge.Age to be -1, got %d", invalidAgeErr.Age)
	}

	// Test error string format
	expectedStr := "age -1 is invalid"
	if err.Error() != expectedStr {
		t.Errorf("Expected error message %q, got %q", expectedStr, err.Error())
	}

	// Test invalid high age
	err = VerifyAge(151)
	if err == nil {
		t.Fatal("VerifyAge(151) expected error, got nil")
	}
	if !errors.As(err, &invalidAgeErr) {
		t.Errorf("Expected error to be of type *ErrInvalidAge, got %T", err)
	} else if invalidAgeErr.Age != 151 {
		t.Errorf("Expected ErrInvalidAge.Age to be 151, got %d", invalidAgeErr.Age)
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

	expected := "Error: age -5 is invalid"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
