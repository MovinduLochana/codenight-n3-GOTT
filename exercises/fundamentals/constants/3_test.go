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

func TestConstants3(t *testing.T) {
	if got := captureOutput(main); got != "43.98226\n" {
		t.Errorf("output = %q; want %q", got, "43.98226\n")
	}
	assertSource(t, "3.go", `(?s)const\s*\([^)]*Pi\s*=\s*3\.14159[^)]*Radius\s*=\s*7\.0`)
	assertSource(t, "3.go", `circumference\s*:=\s*2\s*\*\s*Pi\s*\*\s*Radius`)
}