package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestFindMinMax(t *testing.T) {
	tests := []struct {
		name        string
		nums        []int
		wantMin     int
		wantMax     int
		expectError bool
	}{
		{"normal", []int{12, 45, 2, 8, 99, 14}, 2, 99, false},
		{"single element", []int{5}, 5, 5, false},
		{"all negative", []int{-5, -10, -2, -15}, -15, -2, false},
		{"empty slice", []int{}, 0, 0, true},
		{"nil slice", nil, 0, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			min, max, err := FindMinMax(tt.nums)
			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error for slice %v, got nil", tt.nums)
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, got: %v", err)
				}
				if min != tt.wantMin || max != tt.wantMax {
					t.Errorf("FindMinMax(%v) = (%d, %d); want (%d, %d)", tt.nums, min, max, tt.wantMin, tt.wantMax)
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

	expected := "Min: 2, Max: 99"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
