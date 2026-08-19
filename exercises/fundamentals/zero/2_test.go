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

func TestZero2(t *testing.T) {
	if got := captureOutput(main); got != "0 [] false\n" {
		t.Errorf("output = %q; want %q", got, "0 [] false\n")
	}
	assertSource(t, "2.go", `count\s+int`)
	assertSource(t, "2.go", `message\s+string`)
	assertSource(t, "2.go", `passed\s+bool`)
}