package main

import (
	"errors"
	"strconv"
	"testing"
)

func TestSentinelErrNotPositive(t *testing.T) {
	n, err := ParsePositive("7")
	if err != nil {
		t.Errorf("ParsePositive(\"7\") error = %v; want nil", err)
	}
	if n != 7 {
		t.Errorf("ParsePositive(\"7\") = %d; want 7", n)
	}

	_, err = ParsePositive("-5")
	if !errors.Is(err, ErrNotPositive) {
		t.Errorf("ParsePositive(\"-5\") error must equal ErrNotPositive, got %v", err)
	}

	_, err = ParsePositive("0")
	if !errors.Is(err, ErrNotPositive) {
		t.Errorf("ParsePositive(\"0\") error must equal ErrNotPositive, got %v", err)
	}

	_, err = ParsePositive("abc")
	if err == nil {
		t.Fatalf("ParsePositive(\"abc\") returned nil error; want parse error")
	}
	if errors.Is(err, ErrNotPositive) {
		t.Errorf("ParsePositive(\"abc\") error must NOT be ErrNotPositive, got %v", err)
	}
	if !errors.Is(err, strconv.ErrSyntax) {
		t.Errorf("ParsePositive(\"abc\") error must be the strconv parse error, got %v", err)
	}
}