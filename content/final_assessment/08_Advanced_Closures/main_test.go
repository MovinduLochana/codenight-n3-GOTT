package main

import (
	"bytes"
	"io"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestMakeFilter(t *testing.T) {
	nums := []int{10, 5, 20, 15, 30, 2}

	// Test with limit 10
	filterAbove10 := MakeFilter(10)
	got10 := filterAbove10(nums)
	want10 := []int{20, 15, 30}
	if len(got10) == 0 && len(want10) == 0 {
		// handle empty vs nil nicely, though they should match
	} else if !reflect.DeepEqual(got10, want10) {
		t.Errorf("MakeFilter(10)(%v) = %v; want %v", nums, got10, want10)
	}

	// Test with limit 20
	filterAbove20 := MakeFilter(20)
	got20 := filterAbove20(nums)
	want20 := []int{30}
	if !reflect.DeepEqual(got20, want20) {
		t.Errorf("MakeFilter(20)(%v) = %v; want %v", nums, got20, want20)
	}

	// Test with limit 5
	filterAbove5 := MakeFilter(5)
	got5 := filterAbove5(nums)
	want5 := []int{10, 20, 15, 30}
	if !reflect.DeepEqual(got5, want5) {
		t.Errorf("MakeFilter(5)(%v) = %v; want %v", nums, got5, want5)
	}

	// Test that it does not mutate the original slice
	originalCopy := []int{10, 5, 20, 15, 30, 2}
	if !reflect.DeepEqual(nums, originalCopy) {
		t.Errorf("MakeFilter mutated the original slice. Original was %v, now is %v", originalCopy, nums)
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
		"Above 10: [20 15 30]",
		"Above 20: [30]",
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
