package main

import (
	"errors"
	"strconv"
	"testing"
)

func TestParsePositive(t *testing.T) {
	n, err := ParsePositive("42")
	if err != nil {
		t.Errorf("ParsePositive(\"42\") error = %v; want nil", err)
	}
	if n != 42 {
		t.Errorf("ParsePositive(\"42\") = %d; want 42", n)
	}

	if _, err := ParsePositive("-5"); err == nil {
		t.Errorf("ParsePositive(\"-5\") returned nil error; want non-nil")
	} else if err.Error() != "number must be positive" {
		t.Errorf("ParsePositive(\"-5\") error = %q; want %q", err.Error(), "number must be positive")
	}

	_, err = ParsePositive("abc")
	if err == nil {
		t.Errorf("ParsePositive(\"abc\") returned nil error; want parse error")
	}
	if !errors.Is(err, strconv.ErrSyntax) {
		t.Errorf("ParsePositive(\"abc\") error must be the unwrapped strconv parse error, got %v", err)
	}
}