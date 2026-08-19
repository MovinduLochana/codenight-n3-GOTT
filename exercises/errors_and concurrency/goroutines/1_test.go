package main

import (
	"io"
	"os"
	"testing"
)

func TestGoroutineWG(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	os.Stdout = old
	out, _ := io.ReadAll(r)

	if got := string(out); got != "[0 1 2]\n" {
		t.Errorf("main() output = %q; want %q", got, "[0 1 2]\n")
	}
}