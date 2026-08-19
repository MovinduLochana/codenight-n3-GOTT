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

func TestBasicTypes3(t *testing.T) {
	if got := captureOutput(main); got != "Welcome back Gopher 7 true 4.75\n" {
		t.Errorf("output = %q; want %q", got, "Welcome back Gopher 7 true 4.75\n")
	}
	assertSource(t, "3.go", `name\s*:=\s*"Gopher"`)
	assertSource(t, "3.go", `visits\s*:=\s*7`)
	assertSource(t, "3.go", `rating\s*:=\s*4\.75`)
	assertSource(t, "3.go", `renewed\s*:=\s*visits\s*>\s*3`)
}