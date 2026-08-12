package main

import (
	"bytes"
	"io"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestCharFrequency(t *testing.T) {
	tests := []struct {
		input string
		want  map[rune]int
	}{
		{
			input: "golang programming",
			want: map[rune]int{
				'g': 3, 'o': 2, 'l': 1, 'a': 2, 'n': 2, 'p': 1, 'r': 2, 'm': 2, 'i': 1,
			},
		},
		{
			input: "hello",
			want: map[rune]int{
				'h': 1, 'e': 1, 'l': 2, 'o': 1,
			},
		},
		{
			input: "   ",
			want:  map[rune]int{},
		},
		{
			input: "",
			want:  map[rune]int{},
		},
	}

	for _, tt := range tests {
		got := CharFrequency(tt.input)
		// reflect.DeepEqual can fail if one map is nil and other is empty map.
		// So handle empty map check nicely.
		if len(got) == 0 && len(tt.want) == 0 {
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("CharFrequency(%q) = %v; want %v", tt.input, got, tt.want)
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
		"g: 3",
		"o: 2",
		"l: 1",
		"space: 0",
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
