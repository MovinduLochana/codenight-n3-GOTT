package main

import (
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

func captureOutput(fn func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	fn()
	w.Close()
	os.Stdout = old
	out, _ := io.ReadAll(r)
	return string(out)
}

func assertSource(t *testing.T, file, pattern string) {
	t.Helper()
	data, err := os.ReadFile(file)
	if err != nil {
		t.Fatalf("read %s: %v", file, err)
	}
	if !regexp.MustCompile(pattern).Match(data) {
		t.Errorf("%s does not contain required pattern %q", file, pattern)
	}
}

func TestBasicTypes1(t *testing.T) {
	out := strings.TrimSuffix(captureOutput(main), "\n")
	fields := strings.Split(out, " ")
	if len(fields) != 4 {
		t.Fatalf("expected 4 space-separated values, got %q", out)
	}
	if _, err := strconv.Atoi(fields[0]); err != nil {
		t.Errorf("first value %q should be an int", fields[0])
	}
	if _, err := strconv.ParseFloat(fields[1], 64); err != nil {
		t.Errorf("second value %q should be a float64", fields[1])
	}
	if _, err := strconv.Atoi(fields[3]); err != nil {
		if fields[3] != "true" && fields[3] != "false" {
			t.Errorf("fourth value %q should be a bool", fields[3])
		}
	}
	assertSource(t, "1.go", `age\s+int`)
	assertSource(t, "1.go", `price\s+float64`)
	assertSource(t, "1.go", `city\s+string`)
	assertSource(t, "1.go", `isOpen\s+bool`)
}