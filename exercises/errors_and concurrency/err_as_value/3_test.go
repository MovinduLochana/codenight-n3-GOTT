package main

import (
	"errors"
	"testing"
)

func TestSentinelErrNotPositive(t *testing.T) {
	_, err := ParsePositive("-1")
	if !errors.Is(err, ErrNotPositive) {
		t.Errorf("Expected ErrNotPositive, got %v", err)
	}
}
