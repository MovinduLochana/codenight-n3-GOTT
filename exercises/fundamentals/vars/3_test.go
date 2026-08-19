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

func TestVars3(t *testing.T) {
	if got := captureOutput(main); got != "Golang 2009 true\n" {
		t.Errorf("output = %q; want %q", got, "Golang 2009 true\n")
	}
	assertSource(t, "3.go", `var\s+language\s*=\s*"Golang"`)
	assertSource(t, "3.go", `year\s+int\s*=\s*2009`)
	assertSource(t, "3.go", `isPublic\s+bool\s*=\s*true`)
	data, err := os.ReadFile("3.go")
	if err != nil {
		t.Fatalf("read 3.go: %v", err)
	}
	if regexp.MustCompile(`(?s)func main\(\)\{.*var language`).Match(data) {
		t.Errorf("language must be declared at package level, outside main")
	}
}