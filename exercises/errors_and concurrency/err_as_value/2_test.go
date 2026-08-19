package main

import (
	"errors"
	"strconv"
	"strings"
	"testing"
)

func TestReadAge(t *testing.T) {
	n, err := ReadAge("25")
	if err != nil {
		t.Errorf("ReadAge(\"25\") error = %v; want nil", err)
	}
	if n != 25 {
		t.Errorf("ReadAge(\"25\") = %d; want 25", n)
	}

	_, err = ReadAge("xx")
	if err == nil {
		t.Fatalf("ReadAge(\"xx\") returned nil error; want non-nil")
	}
	if !errors.Is(err, strconv.ErrSyntax) {
		t.Errorf("ReadAge(\"xx\") error must wrap strconv.ErrSyntax, got %v", err)
	}
	if !strings.Contains(err.Error(), "invalid age") {
		t.Errorf("ReadAge(\"xx\") error = %q; want it to contain %q", err.Error(), "invalid age")
	}
}