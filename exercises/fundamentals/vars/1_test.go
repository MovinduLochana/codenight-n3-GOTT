package main

import (
	"io"
	"os"
	"regexp"
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

func TestVars1(t *testing.T) {
	if got := captureOutput(main); got != "Colombo\n" {
		t.Errorf("output = %q; want %q", got, "Colombo\n")
	}
	assertSource(t, "1.go", `var\s+city\s+string`)
	assertSource(t, "1.go", `"Colombo"`)
}